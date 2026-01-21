package service

import (
	"context"
	"net/http"
	"strconv"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/common/http/response"
	"go-link/identity/internal/constant"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/pkg/utils"
	"go-link/identity/internal/ports"
)

type userService struct {
	userRepo         ports.UserRepository
	credentialRepo   ports.CredentialRepository
	tenantRepo       ports.TenantRepository
	tenantMemberRepo ports.TenantMemberRepository
	roleRepo         ports.RoleRepository
	attrDefRepo      ports.AttributeDefinitionRepository
	attrValueRepo    ports.UserAttributeValueRepository
	cache            cache.LocalCache[string, any]
}

// NewUserService creates a new UserService instance.
func NewUserService(
	userRepo ports.UserRepository,
	credentialRepo ports.CredentialRepository,
	tenantRepo ports.TenantRepository,
	tenantMemberRepo ports.TenantMemberRepository,
	roleRepo ports.RoleRepository,
	attrDefRepo ports.AttributeDefinitionRepository,
	attrValueRepo ports.UserAttributeValueRepository,
	cache cache.LocalCache[string, any],
) ports.UserService {
	return &userService{
		userRepo:         userRepo,
		credentialRepo:   credentialRepo,
		tenantRepo:       tenantRepo,
		tenantMemberRepo: tenantMemberRepo,
		roleRepo:         roleRepo,
		attrDefRepo:      attrDefRepo,
		attrValueRepo:    attrValueRepo,
		cache:            cache,
	}
}

// DeleteUser deletes a user by ID
func (s *userService) Delete(ctx context.Context, id int) error {
	exists, err := s.userRepo.Exists(ctx, id)
	if err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to check user exists", http.StatusInternalServerError)
	}

	if !exists {
		return apperr.New(response.CodeNotFound, "user not found", http.StatusNotFound, nil)
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to delete user", http.StatusInternalServerError)
	}

	return nil
}

// UpdateProfile updates user profile
func (s *userService) UpdateProfile(ctx context.Context, userID int, req *dto.UpdateProfileRequest) (*dto.ProfileResponse, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get user", http.StatusInternalServerError)
	}

	attrUpdates := map[string]string{
		constant.AttributeKeyFirstName: req.FirstName,
		constant.AttributeKeyLastName:  req.LastName,
		constant.AttributeKeyGender:    strconv.Itoa(req.Gender),
		constant.AttributeKeyBirthday:  req.Birthday,
	}

	existingAttrs, err := s.attrValueRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get user attributes", http.StatusInternalServerError)
	}

	existingAttrMap := make(map[int]*entity.UserAttributeValue)
	for _, attr := range existingAttrs {
		existingAttrMap[attr.AttributeID] = attr
	}

	var newAttrs []*entity.UserAttributeValue

	for key, value := range attrUpdates {
		if value == "" {
			continue
		}

		def, err := utils.GetAttributeDefinition(ctx, key, s.attrDefRepo, s.cache)
		if err != nil {
			return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get attribute definition: "+key, http.StatusInternalServerError)
		}

		if existingAttr, ok := existingAttrMap[def.ID]; ok {
			// Update existing
			if existingAttr.Value != value {
				existingAttr.Value = value
				if err := s.attrValueRepo.Update(ctx, existingAttr); err != nil {
					return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to update attribute: "+key, http.StatusInternalServerError)
				}
			}
		} else {
			// Create new
			newAttrs = append(newAttrs, &entity.UserAttributeValue{
				UserID:      user.ID,
				AttributeID: def.ID,
				Value:       value,
			})
		}
	}

	if len(newAttrs) > 0 {
		if err := s.attrValueRepo.CreateBulk(ctx, newAttrs); err != nil {
			return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create user attributes", http.StatusInternalServerError)
		}
	}

	gender, _ := strconv.Atoi(attrUpdates[constant.AttributeKeyGender])

	return &dto.ProfileResponse{
		Username:  user.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    gender,
		Birthday:  req.Birthday,
	}, nil
}

// GetProfile gets user profile
func (s *userService) GetProfile(ctx context.Context, userID int) (*dto.ProfileResponse, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get user", http.StatusInternalServerError)
	}

	existingAttrs, err := s.attrValueRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get user attributes", http.StatusInternalServerError)
	}

	// Helper to get value
	firstNameVal := utils.GetAttributeValue(ctx, constant.AttributeKeyFirstName, existingAttrs, s.attrDefRepo, s.cache)
	lastNameVal := utils.GetAttributeValue(ctx, constant.AttributeKeyLastName, existingAttrs, s.attrDefRepo, s.cache)
	genderVal := utils.GetAttributeValue(ctx, constant.AttributeKeyGender, existingAttrs, s.attrDefRepo, s.cache)
	birthdayVal := utils.GetAttributeValue(ctx, constant.AttributeKeyBirthday, existingAttrs, s.attrDefRepo, s.cache)

	genderInt, _ := strconv.Atoi(genderVal)

	return &dto.ProfileResponse{
		Username:  user.Username,
		FirstName: firstNameVal,
		LastName:  lastNameVal,
		Gender:    genderInt,
		Birthday:  birthdayVal,
	}, nil
}
