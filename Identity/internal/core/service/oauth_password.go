package service

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/huynhanx03/GoLink/events-contract/topics"
	"go.uber.org/zap"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/security"

	"go-link/identity/global"
	"go-link/identity/internal/constant"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
	"go-link/identity/pkg/oauth"
)

const (
	oauthTokenType          = "oauth-temp"
	resetTokenType          = "reset-password"
	credentialKeyEmail      = "email"
	credentialKeyExternalID = "external_id"
	credentialKeyVerified   = "verified"
)

// oauthTempClaims holds temporary user info from OAuth providers before a local account is created.
// It allows for a stateless registration flow by carrying verified OAuth data in a signed JWT.
type oauthTempClaims struct {
	jwt.RegisteredClaims
	Email      string `json:"email"`
	ExternalID string `json:"external_id"`
	Provider   string `json:"provider"`
	TokenType  string `json:"type"`
}

// resetClaims contains the user identity for the password reset flow.
// Using a JWT ensures the flow is stateless and secure via RSA digital signatures.
type resetClaims struct {
	jwt.RegisteredClaims
	UserID    int    `json:"user_id"`
	TokenType string `json:"type"`
}

// getProvider resolves the appropriate OAuth handler (e.g., Google, GitHub) based on the provider name.
func (s *authenticationService) getProvider(name string) (oauth.Provider, error) {
	p, ok := s.oauthProviders[name]
	if !ok {
		return nil, apperr.NewError(authServiceName, response.CodeBadRequest, fmt.Sprintf("unsupported oauth provider: %s", name), http.StatusBadRequest, nil)
	}
	return p, nil
}

// OAuthCallback handles the redirect from OAuth providers.
// It either logs in existing users or issues a temporary token to initiate the registration flow for new users.
func (s *authenticationService) OAuthCallback(ctx context.Context, req *dto.OAuthCallbackRequest) (*dto.OAuthCallbackResponse, error) {
	provider, err := s.getProvider(req.Provider)
	if err != nil {
		return nil, err
	}

	// Exchange the authorization code for verified user information from the provider.
	userInfo, err := provider.ExchangeCode(ctx, req.Code)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeBadRequest, "failed to exchange oauth code", http.StatusBadRequest)
	}

	// Check if this external identity is already linked to an existing GoLink user.
	fedIdentity, err := s.fedIdentityRepo.GetByProviderAndExternalID(ctx, req.Provider, userInfo.ExternalID)
	if err == nil && fedIdentity != nil {
		// Existing user found: Generate standard login tokens and complete the flow.
		user, err := s.userRepo.Get(ctx, fedIdentity.UserID)
		if err != nil {
			return nil, err
		}

		refreshToken, err := s.generateRefreshToken(user)
		if err != nil {
			return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
		}

		accessToken, err := s.generateAccessToken(ctx, user, EmptyTenantID, EmptyTierID, nil)
		if err != nil {
			return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
		}

		return &dto.OAuthCallbackResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	// New user: Issue a short-lived 'registration ticket' containing verified OAuth data.
	email, _ := userInfo.Metadata[credentialKeyEmail].(string)
	tempToken, err := s.generateOAuthTempToken(req.Provider, email, userInfo.ExternalID)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
	}

	return &dto.OAuthCallbackResponse{
		RequiresRegistration: true,
		OAuthToken:           tempToken,
	}, nil
}

// OAuthRegister links verified OAuth data with user-chosen credentials (Username/Password).
// This ensures every user has a consistent internal identity regardless of the sign-in method.
func (s *authenticationService) OAuthRegister(ctx context.Context, req *dto.OAuthRegisterRequest) (*dto.LoginResponse, error) {
	// Verify and decode the temporary registration ticket issued in the callback phase.
	claims, err := s.parseOAuthTempToken(req.OAuthToken)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeBadRequest, constant.MsgInvalidGoogleToken, http.StatusBadRequest)
	}

	// Guard against cross-provider token injection.
	if req.Provider != "" && req.Provider != claims.Provider {
		return nil, apperr.NewError(authServiceName, response.CodeBadRequest, "provider mismatch", http.StatusBadRequest, nil)
	}

	var user *entity.User

	// Transactionally create the user profile, personal tenant, and link the external identity.
	err = global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		var err error
		user, err = s.registerInternal(ctx, &dto.CreateUserRequest{
			Username:  req.Username,
			Password:  req.Password,
			IsAdmin:   false,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    req.Gender,
			Birthday:  req.Birthday,
		})
		if err != nil {
			return err
		}

		// Store OAuth email in CredentialData for future account recovery (Forgot Password).
		oauthCred := &entity.Credential{
			UserID: user.ID,
			Type:   claims.Provider,
			CredentialData: map[string]any{
				credentialKeyEmail:      claims.Email,
				credentialKeyExternalID: claims.ExternalID,
				credentialKeyVerified:   true,
			},
		}
		if err := s.credentialRepo.Create(ctx, oauthCred); err != nil {
			return err
		}

		// Permanent link between the local user and the external provider.
		fedIdentity := &entity.FederatedIdentity{
			UserID:     user.ID,
			Provider:   claims.Provider,
			ExternalID: claims.ExternalID,
		}
		return s.fedIdentityRepo.Create(ctx, fedIdentity)
	})

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
	}

	accessToken, err := s.generateAccessToken(ctx, user, EmptyTenantID, EmptyTierID, nil)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// LinkOAuth allows an authenticated user to connect additional OAuth accounts to their GoLink profile.
func (s *authenticationService) LinkOAuth(ctx context.Context, userID int, req *dto.OAuthLinkRequest) (*dto.OAuthLinkResponse, error) {
	provider, err := s.getProvider(req.Provider)
	if err != nil {
		return nil, err
	}

	userInfo, err := provider.ExchangeCode(ctx, req.Code)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeBadRequest, "failed to exchange oauth code", http.StatusBadRequest)
	}

	// Prevent a single external account from being linked to multiple GoLink users.
	existing, _ := s.fedIdentityRepo.GetByProviderAndExternalID(ctx, req.Provider, userInfo.ExternalID)
	if existing != nil {
		return nil, apperr.NewError(authServiceName, response.CodeConflict, constant.MsgGoogleAlreadyUsed, http.StatusConflict, nil)
	}

	err = global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		// Merge all provider metadata into the credential's JSON data for maximum flexibility.
		credData := make(map[string]any)
		for k, v := range userInfo.Metadata {
			credData[k] = v
		}

		credData[credentialKeyExternalID] = userInfo.ExternalID
		if _, ok := credData[credentialKeyVerified]; !ok {
			credData[credentialKeyVerified] = true
		}

		oauthCred := &entity.Credential{
			UserID:         userID,
			Type:           req.Provider,
			CredentialData: credData,
		}
		if err := s.credentialRepo.Create(ctx, oauthCred); err != nil {
			return err
		}

		fedIdentity := &entity.FederatedIdentity{
			UserID:     userID,
			Provider:   req.Provider,
			ExternalID: userInfo.ExternalID,
		}
		return s.fedIdentityRepo.Create(ctx, fedIdentity)
	})

	if err != nil {
		return nil, err
	}

	return &dto.OAuthLinkResponse{Success: true}, nil
}

// ForgotPassword initiates the recovery flow by searching for a verified email within linked OAuth accounts.
func (s *authenticationService) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, error) {
	// Consistently return success to prevent user enumeration attacks.
	successResp := &dto.ForgotPasswordResponse{Message: constant.MsgForgotPasswordMsg}

	// Rate limiting: 1 request per minute per username via Redis
	rateLimitKey := constant.CacheKeyAuthRateLimitForgot + req.Username
	_, found, err := global.Redis.Get(ctx, rateLimitKey)
	if err == nil && found {
		return nil, apperr.NewError(authServiceName, response.CodeBadRequest, constant.MsgRateLimitForgot, http.StatusBadRequest, nil)
	}

	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return successResp, nil
	}

	// Search all registered OAuth credentials for an email address to send the reset token.
	email := s.findUserEmail(ctx, user.ID)
	if email == "" {
		return successResp, nil
	}

	resetToken, err := s.generateResetToken(user.ID)
	if err != nil {
		return successResp, nil
	}

	// Record the request in Redis to trigger rate limiting
	_ = global.Redis.Set(ctx, rateLimitKey, "1", constant.ForgotRateLimitTTL)

	// Dispatch the reset token via the Notification Service through Kafka.
	if s.producer != nil {
		evt := map[string]any{
			"idempotency_key": fmt.Sprintf("forgot-password:%d:%s", user.ID, resetToken[:8]),
			"type":            "forgot-password-otp",
			"channel":         "email",
			"priority":        "urgent",
			"recipient": map[string]any{
				"user_id": strconv.Itoa(user.ID),
				"email":   email,
				"name":    user.Username,
			},
			"template_data": map[string]any{
				"ResetToken": resetToken,
				"Name":       user.Username,
			},
		}

		evtBytes, err := json.Marshal(evt)
		if err == nil {
			if _, _, pubErr := s.producer.Publish(ctx, topics.NotificationSend, []byte(strconv.Itoa(user.ID)), evtBytes); pubErr != nil {
				global.LoggerZap.Error("Failed to publish notification event", zap.Error(pubErr))
			}
		}
	}

	return successResp, nil
}

// findUserEmail scans all linked OAuth credentials to find a verified email address for account recovery.
func (s *authenticationService) findUserEmail(ctx context.Context, userID int) string {
	for providerName := range s.oauthProviders {
		cred, err := s.credentialRepo.GetByUserID(ctx, userID, providerName)
		if err != nil || cred == nil {
			continue
		}

		email, _ := cred.CredentialData[credentialKeyEmail].(string)
		if email != "" {
			return email
		}
	}
	return ""
}

// ResetPassword validates the reset token and updates the local password credential.
func (s *authenticationService) ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error) {
	// Verify the JWT's signature and expiration.
	claims, err := s.parseResetToken(req.Token)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeBadRequest, constant.MsgInvalidResetToken, http.StatusBadRequest)
	}

	// Check if the token has already been used (Single-use policy)
	blacklistKey := constant.CacheKeyAuthBlacklistJTI + claims.ID
	_, found, err := global.Redis.Get(ctx, blacklistKey)
	if err == nil && found {
		return nil, apperr.NewError(authServiceName, response.CodeBadRequest, constant.MsgTokenAlreadyUsed, http.StatusBadRequest, nil)
	}

	newHash, err := security.HashPassword(req.NewPassword)
	if err != nil {
		return nil, apperr.MapError(authServiceName, err, response.CodeInternalError, apperr.MsgGenFailed, http.StatusInternalServerError)
	}

	err = global.EntClient.DoInTx(ctx, func(ctx context.Context) error {
		cred, err := s.credentialRepo.GetByUserID(ctx, claims.UserID, credentialTypePassword)
		if err != nil {
			return err
		}
		// Update the 'hash' field within the password credential's JSON data.
		cred.CredentialData[credentialKeyHash] = newHash
		if err := s.credentialRepo.Update(ctx, cred); err != nil {
			return err
		}

		// Blacklist the token's JTI in Redis once successfully used.
		return global.Redis.Set(ctx, blacklistKey, "1", constant.ResetTokenTTL)
	})

	if err != nil {
		return nil, err
	}

	return &dto.ResetPasswordResponse{Success: true}, nil
}

// generateOAuthTempToken builds a signed JWT to carry OAuth data during the registration process.
func (s *authenticationService) generateOAuthTempToken(provider, email, externalID string) (string, error) {
	privateKey := global.Config.JWT.PrivateKey
	claims := &oauthTempClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.OAuthTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
		Email:      email,
		ExternalID: externalID,
		Provider:   provider,
		TokenType:  oauthTokenType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// parseOAuthTempToken parses the oauth temp token with the given public key.
func (s *authenticationService) parseOAuthTempToken(tokenStr string) (*oauthTempClaims, error) {
	publicKey := global.Config.JWT.PublicKey
	claims := &oauthTempClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid oauth temp token: %w", err)
	}
	if claims.TokenType != oauthTokenType {
		return nil, fmt.Errorf("wrong token type: %s", claims.TokenType)
	}

	return claims, nil
}

// generateResetToken generates a signed JWT to carry reset token data during the registration process.
func (s *authenticationService) generateResetToken(userID int) (string, error) {
	privateKey := global.Config.JWT.PrivateKey
	claims := &resetClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.ResetTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
		UserID:    userID,
		TokenType: resetTokenType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// parseResetToken parses the reset token with the given public key.
func (s *authenticationService) parseResetToken(tokenStr string) (*resetClaims, error) {
	publicKey := global.Config.JWT.PublicKey
	return parseResetTokenWithKey(tokenStr, publicKey)
}

// parseResetTokenWithKey parses the reset token with the given public key.
func parseResetTokenWithKey(tokenStr string, publicKey *rsa.PublicKey) (*resetClaims, error) {
	claims := &resetClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid reset token: %w", err)
	}
	if claims.TokenType != resetTokenType {
		return nil, fmt.Errorf("wrong token type: %s", claims.TokenType)
	}

	return claims, nil
}
