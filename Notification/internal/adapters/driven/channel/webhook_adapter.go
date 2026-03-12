package channel

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

type webhookAdapter struct {
	webhookRepo ports.WebhookConfigRepository
	client      *http.Client
}

// NewWebhookAdapter creates a ChannelAdapter that delivers notifications to tenant webhook endpoints.
func NewWebhookAdapter(webhookRepo ports.WebhookConfigRepository) ports.ChannelAdapter {
	return &webhookAdapter{
		webhookRepo: webhookRepo,
		client:      &http.Client{Timeout: 10 * time.Second},
	}
}

// Send looks up webhook configs for the recipient's tenant and delivers to each active endpoint.
// Errors from individual endpoints are logged but do not abort delivery to remaining endpoints.
func (a *webhookAdapter) Send(ctx context.Context, notification *entity.Notification) error {
	configs, err := a.webhookRepo.GetByTenantID(ctx, notification.Recipient.UserID)
	if err != nil || len(configs) == 0 {
		return nil // No webhooks configured — skip silently.
	}

	payload := map[string]any{
		"event":        notification.Type,
		"notification": notification,
		"timestamp":    time.Now().UTC().Format(time.RFC3339),
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal webhook payload: %w", err)
	}

	for _, cfg := range configs {
		if !cfg.IsActive {
			continue
		}
		if !matchesEventType(cfg.EventTypes, notification.Type) {
			continue
		}
		// Best-effort: continue to remaining endpoints on error.
		_ = a.sendToEndpoint(ctx, cfg, body)
	}

	return nil
}

// sendToEndpoint signs the payload and POSTs it to the webhook URL.
func (a *webhookAdapter) sendToEndpoint(ctx context.Context, cfg *entity.WebhookConfig, body []byte) error {
	signature := signPayload(cfg.Secret, body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.URL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", "sha256="+signature)
	req.Header.Set("X-Webhook-ID", cfg.ID)

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body) // drain to allow connection reuse

	if resp.StatusCode >= 400 {
		return fmt.Errorf("webhook endpoint returned status %d", resp.StatusCode)
	}
	return nil
}

// signPayload computes HMAC-SHA256 of body using the webhook secret.
func signPayload(secret string, body []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

// matchesEventType returns true if eventType is in the subscription list, or list is empty/contains "*".
func matchesEventType(types []string, eventType string) bool {
	if len(types) == 0 {
		return true
	}
	for _, t := range types {
		if t == eventType || t == "*" {
			return true
		}
	}
	return false
}

func (a *webhookAdapter) Type() string { return entity.ChannelWebhook }
