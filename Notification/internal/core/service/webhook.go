package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/notification/internal/core/dto"
	"go-link/notification/internal/core/mapper"
	"go-link/notification/internal/ports"
)

type webhookConfigService struct {
	repo ports.WebhookConfigRepository
}

// NewWebhookConfigService creates a new WebhookConfigService.
func NewWebhookConfigService(repo ports.WebhookConfigRepository) ports.WebhookConfigService {
	return &webhookConfigService{repo: repo}
}

// generateSecret produces a cryptographically random 32-byte hex string for HMAC signing.
func generateSecret() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func getTenantID(ctx context.Context) (string, error) {
	tenantID, ok := ctx.Value(constraints.ContextKeyUserID).(string)
	if !ok || tenantID == "" {
		return "", apperr.New(response.CodeUnauthorized, "unauthorized", 0, nil)
	}
	return tenantID, nil
}

func (s *webhookConfigService) Create(ctx context.Context, req *dto.CreateWebhookRequest) (*dto.WebhookResponse, error) {
	tenantID, err := getTenantID(ctx)
	if err != nil {
		return nil, err
	}

	cfg := mapper.ToWebhookEntityFromCreate(tenantID, req)
	secret, err := generateSecret()
	if err != nil {
		return nil, apperr.New(response.CodeInternalServer, "failed to generate secret", 0, err)
	}

	cfg.Secret = secret
	cfg.CreatedAt = time.Now()
	cfg.UpdatedAt = time.Now()

	if err := s.repo.Create(ctx, cfg); err != nil {
		return nil, apperr.New(response.CodeInternalServer, "failed to create webhook config", 0, err)
	}

	return mapper.ToWebhookResponse(cfg), nil
}

func (s *webhookConfigService) Get(ctx context.Context, id string) (*dto.WebhookResponse, error) {
	tenantID, err := getTenantID(ctx)
	if err != nil {
		return nil, err
	}

	cfg, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, apperr.New(response.CodeInternalServer, "webhook not found", 0, err)
	}
	if cfg.TenantID != tenantID {
		return nil, apperr.New(response.CodeUnauthorized, "unauthorized access to webhook", 0, nil)
	}
	return mapper.ToWebhookResponse(cfg), nil
}

func (s *webhookConfigService) FindByTenantID(ctx context.Context) ([]*dto.WebhookResponse, error) {
	tenantID, err := getTenantID(ctx)
	if err != nil {
		return nil, err
	}

	configs, err := s.repo.GetByTenantID(ctx, tenantID)
	if err != nil {
		return nil, apperr.New(response.CodeInternalServer, "failed to list webhooks", 0, err)
	}
	return mapper.ToWebhookResponseList(configs), nil
}

func (s *webhookConfigService) Update(ctx context.Context, req *dto.UpdateWebhookRequest) (*dto.WebhookResponse, error) {
	tenantID, err := getTenantID(ctx)
	if err != nil {
		return nil, err
	}

	// Verify ownership and existence
	oldCfg, err := s.repo.Get(ctx, req.ID)
	if err != nil {
		return nil, apperr.New(response.CodeInternalServer, "webhook not found", 0, err)
	}
	if oldCfg.TenantID != tenantID {
		return nil, apperr.New(response.CodeUnauthorized, "unauthorized access to webhook", 0, nil)
	}

	cfg := mapper.ToWebhookEntityFromUpdate(tenantID, req.ID, req)
	// Maintain secret and creation date
	cfg.Secret = oldCfg.Secret
	cfg.CreatedAt = oldCfg.CreatedAt
	cfg.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, cfg); err != nil {
		return nil, apperr.New(response.CodeInternalServer, "failed to update webhook", 0, err)
	}

	return mapper.ToWebhookResponse(cfg), nil
}

func (s *webhookConfigService) Delete(ctx context.Context, id string) error {
	tenantID, err := getTenantID(ctx)
	if err != nil {
		return err
	}

	cfg, err := s.repo.Get(ctx, id)
	if err != nil {
		return apperr.New(response.CodeInternalServer, "webhook not found", 0, err)
	}
	if cfg.TenantID != tenantID {
		return apperr.New(response.CodeUnauthorized, "unauthorized access to webhook", 0, nil)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return apperr.New(response.CodeInternalServer, "failed to delete webhook", 0, err)
	}

	return nil
}
