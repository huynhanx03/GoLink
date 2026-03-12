package http

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"

	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/constraints"
	"go-link/notification/internal/core/service"
)

// SSEHandler handles the real-time Server-Sent Events stream endpoint.
type SSEHandler struct {
	hub *service.SSEHub
}

// NewSSEHandler creates a new SSEHandler.
func NewSSEHandler(hub *service.SSEHub) *SSEHandler {
	return &SSEHandler{hub: hub}
}

// Stream handles GET /notifications/stream
// Registers the authenticated user as an SSE client and streams notification
// events until the client disconnects.
func (h *SSEHandler) Stream(c *gin.Context) {
	userID, ok := c.Request.Context().Value(constraints.ContextKeyUserID).(string)
	if !ok || userID == "" {
		response.ErrorResponse(c, response.CodeUnauthorized, nil)
		return
	}

	// Register this connection in the SSE hub.
	client := h.hub.Register(userID)
	defer h.hub.Unregister(client)

	// Set SSE-specific response headers.
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no") // Disable nginx buffering.

	// Send an initial "connected" heartbeat so the client knows the stream is live.
	fmt.Fprintf(c.Writer, "event: connected\ndata: {}\n\n")
	c.Writer.Flush()

	ctx := c.Request.Context()

	for {
		select {
		case <-ctx.Done():
			// Client disconnected.
			return

		case notification, open := <-client.Chan:
			if !open {
				// Hub unregistered this client — channel closed.
				return
			}

			data, err := json.Marshal(notification)
			if err != nil {
				continue
			}

			_, writeErr := fmt.Fprintf(c.Writer, "event: notification\ndata: %s\n\n", data)
			if writeErr != nil {
				// Client gone (broken pipe etc.).
				if writeErr == io.EOF {
					return
				}
				return
			}

			c.Writer.Flush()
		}
	}
}
