package channel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

const resendAPIURL = "https://api.resend.com/emails"

type resendEmailAdapter struct {
	apiKey    string
	fromEmail string
	fromName  string
	client    *http.Client
}

// NewResendEmailAdapter creates a new email channel adapter backed by the Resend API.
func NewResendEmailAdapter(apiKey, fromEmail, fromName string) ports.ChannelAdapter {
	return &resendEmailAdapter{
		apiKey:    apiKey,
		fromEmail: fromEmail,
		fromName:  fromName,
		client:    &http.Client{Timeout: 10 * time.Second},
	}
}

// resendRequest is the payload sent to the Resend API.
type resendRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
}

// Send delivers a notification via the Resend email API.
func (a *resendEmailAdapter) Send(ctx context.Context, notification *entity.Notification) error {
	from := fmt.Sprintf("%s <%s>", a.fromName, a.fromEmail)
	reqBody := resendRequest{
		From:    from,
		To:      []string{notification.Recipient.Email},
		Subject: notification.Subject,
		HTML:    notification.Body,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal resend request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, resendAPIURL, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("create resend request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("send resend request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // 1 MB max
		return fmt.Errorf("resend API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// Type returns the channel type identifier.
func (a *resendEmailAdapter) Type() string {
	return entity.ChannelEmail
}
