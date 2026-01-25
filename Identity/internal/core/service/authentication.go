package service

import (
	"context"
	"net/http"
	"strconv"

	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/common/http/response"
	p "go-link/common/pkg/permissions"
	"go-link/common/pkg/security"
	"go-link/common/pkg/utils"

	"go-link/identity/global"
	"go-link/identity/internal/constant"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
	u "go-link/identity/pkg/utils"
)

const (
	authServiceName = "AuthenticationService"

	credentialTypePassword = "password"
	roleNameOwner          = "owner"
	defaultTierID          = 1
	credentialKeyHash      = "hash"
)

type authenticationService struct {
	userRepo         ports.UserRepository
	credentialRepo   ports.CredentialRepository
	tenantRepo       ports.TenantRepository
	tenantMemberRepo ports.TenantMemberRepository
	roleRepo         ports.RoleRepository
	permissionRepo   ports.PermissionRepository
	resourceRepo     ports.ResourceRepository
	attrDefRepo      ports.AttributeDefinitionRepository
	attrValueRepo    ports.UserAttributeValueRepository
	cache            cache.LocalCache[string, any]
	cacheService     ports.CacheService
}

// NewAuthenticationService creates a new AuthenticationService instance.
func NewAuthenticationService(
	userRepo ports.UserRepository,
	credentialRepo ports.CredentialRepository,
	tenantRepo ports.TenantRepository,
	tenantMemberRepo ports.TenantMemberRepository,
	roleRepo ports.RoleRepository,
	permissionRepo ports.PermissionRepository,
	resourceRepo ports.ResourceRepository,
	attrDefRepo ports.AttributeDefinitionRepository,
	attrValueRepo ports.UserAttributeValueRepository,
	cache cache.LocalCache[string, any],
	cacheService ports.CacheService,
) ports.AuthenticationService {
	return &authenticationService{
		userRepo:         userRepo,
		credentialRepo:   credentialRepo,
		tenantRepo:       tenantRepo,
		tenantMemberRepo: tenantMemberRepo,
		roleRepo:         roleRepo,
		permissionRepo:   permissionRepo,
		resourceRepo:     resourceRepo,
		attrDefRepo:      attrDefRepo,
		attrValueRepo:    attrValueRepo,
		cache:            cache,
		cacheService:     cacheService,
	}
}

// Register creates a new user with a personal tenant.
func (s *authenticationService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	err := global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		createUserReq := &dto.CreateUserRequest{
			Username:  req.Username,
			Password:  req.Password,
			IsAdmin:   false,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    req.Gender,
			Birthday:  req.Birthday,
		}

		user, err := s.CreateUser(ctx, createUserReq)
		if err != nil {
			return err
		}

		// Create personal tenant with username as name
		tenant := &entity.Tenant{
			Name:   req.Username,
			TierID: defaultTierID, // TODO: UPDATE BILLING SERVICE
		}
		if err := s.tenantRepo.Create(ctx, tenant); err != nil {
			return err
		}

		cacheKey := constant.CacheKeyPrefixRoleName + roleNameOwner
		var ownerRole *entity.Role

		if r, found := cache.GetLocal[*entity.Role](s.cache, cacheKey); found {
			ownerRole = r
		}

		if ownerRole == nil {
			var err error
			ownerRole, err = s.roleRepo.GetByName(ctx, roleNameOwner)
			if err != nil {
				return err
			}
			cache.SetLocal(s.cache, cacheKey, ownerRole, constant.CacheCostRoleName)
		}

		// Assign user as owner of the tenant
		tenantMember := &entity.TenantMember{
			TenantID: tenant.ID,
			UserID:   user.ID,
			RoleID:   ownerRole.ID,
		}
		if err := s.tenantMemberRepo.Create(ctx, tenantMember); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{Success: true}, nil
}

// Login processes user login.
func (s *authenticationService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Find user by username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// Get credential
	cred, err := s.credentialRepo.GetByUserID(ctx, user.ID, credentialTypePassword)
	if err != nil {
		return nil, err
	}

	// Verify password
	hash, ok := cred.CredentialData[credentialKeyHash].(string)
	if !ok {
		return nil, NewError(authServiceName, response.CodeInternalError, MsgInvalidCredData, http.StatusInternalServerError, nil)
	}

	if err := security.ComparePassword(hash, req.Password); err != nil {
		return nil, MapError(authServiceName, err, response.CodeUnauthorized, MsgInvalidAuth, http.StatusUnauthorized)
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return nil, MapError(authServiceName, err, response.CodeInternalError, MsgGenFailed, http.StatusInternalServerError)
	}

	var accessToken string
	if user.IsAdmin {
		accessToken, err = s.generateAccessToken(ctx, user, 0, 0, nil)
		if err != nil {
			return nil, MapError(authServiceName, err, response.CodeInternalError, MsgGenFailed, http.StatusInternalServerError)
		}
	}

	return &dto.LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

// AcquireToken handles tenant token acquisition (Derived Token).
func (s *authenticationService) AcquireToken(ctx context.Context, req *dto.AcquireTokenRequest) (*dto.AcquireTokenResponse, error) {
	// Get tenant member
	member, err := s.tenantMemberRepo.GetByUserAndTenant(ctx, req.UserID, req.TenantID)
	if err != nil {
		return nil, err
	}

	// Get role
	role, err := s.roleRepo.Get(ctx, member.RoleID)
	if err != nil {
		return nil, err
	}

	// Get tenant
	tenant, err := s.tenantRepo.Get(ctx, req.TenantID)
	if err != nil {
		return nil, err
	}

	// Fetch user details for token claims
	user, err := s.userRepo.Get(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.generateAccessToken(ctx, user, req.TenantID, tenant.TierID, role)
	if err != nil {
		return nil, err
	}

	return &dto.AcquireTokenResponse{
		AccessToken: accessToken,
	}, nil
}

// ChangePassword processes password change request.
func (s *authenticationService) ChangePassword(ctx context.Context, userID int, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error) {
	cred, err := s.credentialRepo.GetByUserID(ctx, userID, credentialTypePassword)
	if err != nil {
		return nil, err
	}

	// Verify current password
	hash, ok := cred.CredentialData[credentialKeyHash].(string)
	if !ok {
		return nil, NewError(authServiceName, response.CodeInternalError, MsgInvalidCredData, http.StatusInternalServerError, nil)
	}

	if err := security.ComparePassword(hash, req.CurrentPassword); err != nil {
		return nil, MapError(authServiceName, err, response.CodeUnauthorized, MsgPassIncorrect, http.StatusUnauthorized)
	}

	// Hash new password
	newHash, err := security.HashPassword(req.NewPassword)
	if err != nil {
		return nil, MapError(authServiceName, err, response.CodeInternalError, MsgGenFailed, http.StatusInternalServerError)
	}

	err = global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		// Update credential
		cred.CredentialData[credentialKeyHash] = newHash
		if err := s.credentialRepo.Update(ctx, cred); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &dto.ChangePasswordResponse{Success: true}, nil
}

// CreateUser creates a new user.
func (s *authenticationService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*entity.User, error) {
	exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, NewError(authServiceName, response.CodeConflict, MsgUsernameExists, http.StatusConflict, nil)
	}

	var user *entity.User
	err = global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		user = &entity.User{
			Username: req.Username,
			IsAdmin:  req.IsAdmin,
		}

		if err := s.userRepo.Create(ctx, user); err != nil {
			return err
		}

		// Hash password and create credential
		hashedPassword, err := security.HashPassword(req.Password)
		if err != nil {
			return MapError(authServiceName, err, response.CodeInternalError, MsgGenFailed, http.StatusInternalServerError)
		}

		credential := &entity.Credential{
			UserID: user.ID,
			Type:   credentialTypePassword,
			CredentialData: map[string]any{
				credentialKeyHash: hashedPassword,
			},
		}

		if err := s.credentialRepo.Create(ctx, credential); err != nil {
			return err
		}

		attrKeys := map[string]string{
			constant.AttributeKeyFirstName: req.FirstName,
			constant.AttributeKeyLastName:  req.LastName,
			constant.AttributeKeyGender:    strconv.Itoa(req.Gender),
			constant.AttributeKeyBirthday:  req.Birthday,
		}

		var attrValues []*entity.UserAttributeValue
		for key, value := range attrKeys {
			if value == "" {
				continue
			}

			def, err := u.GetAttributeDefinition(ctx, key, s.attrDefRepo, s.cache)
			if err != nil {
				return err
			}

			attrValues = append(attrValues, &entity.UserAttributeValue{
				UserID:      user.ID,
				AttributeID: def.ID,
				Value:       value,
			})
		}

		if len(attrValues) > 0 {
			if err := s.attrValueRepo.CreateBulk(ctx, attrValues); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

// generateAccessToken generates a new Access Token with aggregated permissions.
func (s *authenticationService) generateAccessToken(ctx context.Context, user *entity.User, tenantID int, tierID int, role *entity.Role) (string, error) {
	secret := global.Config.JWT.Secret

	var permissions map[int]int
	var roleNames []string
	var roleLevel int

	if role != nil {
		roleNames = []string{role.Name}
		roleLevel = role.Level

		permVersion, _ := s.cacheService.GetPermissionConfigVersion(ctx)
		cacheKey := constant.CacheKeyPrefixRolePermissions + strconv.FormatInt(permVersion, 10) + "::" + strconv.Itoa(role.ID)
		if p, found := cache.GetLocal[map[int]int](s.cache, cacheKey); found {
			permissions = p
		}

		if permissions == nil {
			descendants, err := s.roleRepo.FindDescendants(ctx, role.Lft, role.Rgt)
			if err != nil {
				return "", err
			}

			roleIDs := make([]int, len(descendants))
			for i, r := range descendants {
				roleIDs[i] = r.ID
			}

			perms, err := s.permissionRepo.FindByRoleIDs(ctx, roleIDs)
			if err != nil {
				return "", err
			}

			if len(perms) == 0 {
				permissions = make(map[int]int)
			} else {
				permissions = make(map[int]int)
				for _, p := range perms {
					permissions[p.ResourceID] |= p.Scopes
				}
				cache.SetLocal(s.cache, cacheKey, permissions, constant.CacheCostRolePermissions)
			}
		}
	}

	permissionsBlob, err := p.Compress(permissions)
	if err != nil {
		return "", MapError(authServiceName, err, response.CodeInternalError, MsgProcessFailed, http.StatusInternalServerError)
	}

	return utils.GenerateToken(
		secret,
		user.ID,
		user.Username,
		user.IsAdmin,
		tenantID,
		tierID,
		roleNames,
		roleLevel,
		permissionsBlob,
		utils.AccessToken,
	)
}

// generateRefreshToken generates a new Refresh Token.
func (s *authenticationService) generateRefreshToken(user *entity.User) (string, error) {
	secret := global.Config.JWT.Secret
	return utils.GenerateToken(secret, user.ID, user.Username, user.IsAdmin, 0, 0, nil, 0, "", utils.RefreshToken)
}

// RefreshToken refreshes the access token.
func (s *authenticationService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int) (*dto.RefreshTokenResponse, error) {
	// Refresh token for admin
	if req.TenantID == 0 {
		user, err := s.userRepo.Get(ctx, userID)

		if err != nil {
			return nil, err
		}

		if user.IsAdmin == false {
			return nil, NewError(authServiceName, response.CodeUnauthorized, MsgUnauthorized, http.StatusUnauthorized, nil)
		}

		token, err := s.generateAccessToken(ctx, user, 0, 0, nil)
		if err != nil {
			return nil, MapError(authServiceName, err, response.CodeInternalError, MsgGenFailed, http.StatusInternalServerError)
		}

		return &dto.RefreshTokenResponse{
			AccessToken: token,
		}, nil
	}

	request := &dto.AcquireTokenRequest{
		UserID:   userID,
		TenantID: req.TenantID,
	}

	// Refresh token for tenant
	response, err := s.AcquireToken(ctx, request)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: response.AccessToken,
	}, nil
}
