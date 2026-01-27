package service

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	billingv1 "go-link/common/gen/go/billing/v1"
	identityv1 "go-link/common/gen/go/identity/v1"
	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/common/pkg/utils"

	"go-link/generation/global"
	"go-link/generation/internal/constant"
	"go-link/generation/internal/core/dto"
	"go-link/generation/internal/core/entity"
	"go-link/generation/internal/core/mapper"
	"go-link/generation/internal/ports"
)

type linkService struct {
	linkRepo       ports.LinkRepository
	linkCache      ports.LinkCacheRepository
	codePool       ports.ShortCodePool
	localCache     cache.LocalCache[string, int]
	identityClient identityv1.IdentityServiceClient
	billingClient  billingv1.BillingServiceClient
}

func NewLinkService(
	linkRepo ports.LinkRepository,
	codePool ports.ShortCodePool,
	linkCache ports.LinkCacheRepository,
	localCache cache.LocalCache[string, int],
	identityClient identityv1.IdentityServiceClient,
	billingClient billingv1.BillingServiceClient,
) ports.LinkService {
	return &linkService{
		linkRepo:       linkRepo,
		linkCache:      linkCache,
		codePool:       codePool,
		localCache:     localCache,
		identityClient: identityClient,
		billingClient:  billingClient,
	}
}

const serviceName = "LinkService"

// Create creates a new link
// Create creates a new link
func (s *linkService) Create(ctx context.Context, req *dto.CreateLinkRequest) (*dto.LinkResponse, error) {
	link := mapper.ToLinkEntityFromReq(req)
	shortCode := s.codePool.GetOrGenerate()
	link.ID = shortCode

	claims, isUser := ctx.Value(constraints.ContextKeyClaims).(*utils.Claims)
	if !isUser {
		// Guest User
		link.UserID = 0
		link.TenantID = 0
	} else {
		// Authenticated User
		err := s.checkQuota(ctx, claims.TenantID, claims.TierID)
		if err != nil {
			return nil, err
		}
		link.UserID = claims.UserID
		link.TenantID = claims.TenantID
	}

	if err := s.linkRepo.Create(ctx, link, 0); err != nil {
		if isUser {
			s.linkCache.DecrementQuota(ctx, claims.TenantID)
		}
		return nil, apperr.MapError(serviceName, err, response.CodeDatabaseError, apperr.MsgCreateFailed, http.StatusInternalServerError)
	}

	if err := s.linkCache.Set(ctx, link); err != nil {
		// TODO: Log error
	}

	return mapper.ToLinkResponse(link), nil
}

// Delete deletes a link
func (s *linkService) Delete(ctx context.Context, req *dto.DeleteLinkRequest) error {
	link, err := s.linkRepo.Get(ctx, req.ID)
	if err != nil {
		return apperr.NewError(serviceName, response.CodeNotFound, apperr.MsgNotFound, http.StatusNotFound, err)
	}

	userID, _ := ctx.Value(constraints.ContextKeyUserID).(int)
	roleLevel, _ := ctx.Value(constraints.ContextKeyRoleLevel).(int)
	tenantID, _ := ctx.Value(constraints.ContextKeyTenantID).(int)

	if err := s.checkPermission(ctx, link, userID, roleLevel, tenantID); err != nil {
		return err
	}

	if err := s.linkRepo.Delete(ctx, req.ID); err != nil {
		return apperr.NewError(serviceName, response.CodeDatabaseError, apperr.MsgDeleteFailed, http.StatusInternalServerError, err)
	}

	s.linkCache.DecrementQuota(ctx, link.TenantID)

	return nil
}

// checkQuota checks if the user has enough quota to create a link
func (s *linkService) checkQuota(ctx context.Context, tenantID int, tierID int) error {
	cacheKey := fmt.Sprintf(constant.LocalCacheKeyTierConfig, tierID)
	maxQuota, found := s.localCache.Get(cacheKey)

	if !found {
		if s.billingClient == nil {
			return apperr.NewError(serviceName, response.CodeInternalError, constant.MsgBillingUnavailable, http.StatusInternalServerError, nil)
		}

		resp, err := s.billingClient.GetTierConfig(ctx, &billingv1.GetTierConfigRequest{
			TierId: int64(tierID),
		})

		if err != nil {
			global.LoggerZap.Error("Failed to get tier config from Billing", zap.Error(err))
			return apperr.NewError(serviceName, response.CodeInternalError, constant.MsgGetTierConfigFailed, http.StatusInternalServerError, err)
		}

		if resp.MaxLinks == -1 {
			maxQuota = 1_000_000_000 // Unlimited
		} else {
			maxQuota = int(resp.MaxLinks)
		}

		s.localCache.Set(cacheKey, maxQuota, constant.CacheCostQuota)
	}

	usage, err := s.linkCache.IncrementQuota(ctx, tenantID)
	if err != nil {
		global.LoggerZap.Error("Failed to incr usage", zap.Error(err))
		return apperr.NewError(serviceName, response.CodeInternalError, constant.MsgInternalError, http.StatusInternalServerError, err)
	}

	if int(usage) > maxQuota {
		s.linkCache.DecrementQuota(ctx, tenantID)
		return apperr.NewError(serviceName, response.CodeForbidden, constant.MsgQuotaExceeded, http.StatusForbidden, nil)
	}

	return nil
}

// checkPermission checks if the user has permission to delete the link
func (s *linkService) checkPermission(ctx context.Context, link *entity.Link, userID int, roleLevel int, tenantID int) error {
	if link.UserID == userID {
		return nil
	}

	if link.TenantID != tenantID {
		return apperr.NewError(serviceName, response.CodeForbidden, constant.MsgInsufficientPermission, http.StatusForbidden, nil)
	}

	ownerLevel, err := s.getOwnerLevel(ctx, link.UserID, link.TenantID)
	if err != nil {
		global.LoggerZap.Warn("Failed to get owner level", zap.Error(err))
		return apperr.NewError(serviceName, response.CodeInternalError, constant.MsgVerifyPermissionFailed, http.StatusInternalServerError, err)
	}

	if roleLevel <= ownerLevel {
		return apperr.NewError(serviceName, response.CodeForbidden, constant.MsgInsufficientPermission, http.StatusForbidden, nil)
	}

	return nil
}

// getOwnerLevel gets the owner level of the user
func (s *linkService) getOwnerLevel(ctx context.Context, userID int, tenantID int) (int, error) {
	level, err := s.linkCache.GetUserLevel(ctx, userID)
	if err == nil {
		return level, nil
	}

	if s.identityClient == nil {
		return 0, fmt.Errorf("identity client not available")
	}

	resp, err := s.identityClient.GetUserRole(ctx, &identityv1.GetUserRoleRequest{
		UserId:   int64(userID),
		TenantId: int64(tenantID),
	})
	if err != nil {
		return 0, err
	}

	if resp.Role == nil {
		return 0, nil
	}

	ownerLevel := int(resp.Role.Level)
	s.linkCache.SetUserLevel(ctx, userID, ownerLevel)

	return ownerLevel, nil
}
