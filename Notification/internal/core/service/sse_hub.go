package service

import (
	"sync"

	"go-link/notification/internal/core/entity"
)

// SSEClient represents a single SSE connection for a user.
type SSEClient struct {
	UserID string
	Chan   chan *entity.Notification
}

// SSEHub manages all active SSE client connections keyed by user ID.
type SSEHub struct {
	mu      sync.RWMutex
	clients map[string][]*SSEClient
}

// NewSSEHub creates and returns a new SSEHub.
func NewSSEHub() *SSEHub {
	return &SSEHub{
		clients: make(map[string][]*SSEClient),
	}
}

// Register creates a new SSEClient for the given user and registers it in the hub.
// The caller is responsible for calling Unregister when the connection closes.
func (h *SSEHub) Register(userID string) *SSEClient {
	client := &SSEClient{
		UserID: userID,
		Chan:   make(chan *entity.Notification, 16),
	}

	h.mu.Lock()
	h.clients[userID] = append(h.clients[userID], client)
	h.mu.Unlock()

	return client
}

// Unregister removes the given SSEClient from the hub and closes its channel.
func (h *SSEHub) Unregister(client *SSEClient) {
	h.mu.Lock()
	defer h.mu.Unlock()

	conns := h.clients[client.UserID]
	filtered := conns[:0]
	for _, c := range conns {
		if c != client {
			filtered = append(filtered, c)
		}
	}

	if len(filtered) == 0 {
		delete(h.clients, client.UserID)
	} else {
		h.clients[client.UserID] = filtered
	}

	close(client.Chan)
}

// Broadcast sends a notification to all active SSE connections for the given user.
// Non-blocking: drops the event if a client's channel buffer is full.
// Safe against concurrent Unregister: copies the slice under lock.
func (h *SSEHub) Broadcast(userID string, notification *entity.Notification) {
	h.mu.RLock()
	conns := make([]*SSEClient, len(h.clients[userID]))
	copy(conns, h.clients[userID])
	h.mu.RUnlock()

	for _, c := range conns {
		select {
		case c.Chan <- notification:
		default:
			// Client channel full — drop event rather than block.
		}
	}
}
