package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"
	"go-link/notification/global"
	driverHttp "go-link/notification/internal/adapters/driver/http"
)

// RouterGroup contains all handlers decoupled via interfaces.
type RouterGroup struct {
	NotificationHandler  driverHttp.NotificationHandler
	PreferenceHandler    driverHttp.PreferenceHandler
	WebhookConfigHandler driverHttp.WebhookConfigHandler
	SSEHandler           *driverHttp.SSEHandler
}

// NewRouterGroup creates a new RouterGroup.
func NewRouterGroup(
	notificationHandler driverHttp.NotificationHandler,
	preferenceHandler driverHttp.PreferenceHandler,
	webhookConfigHandler driverHttp.WebhookConfigHandler,
	sseHandler *driverHttp.SSEHandler,
) *RouterGroup {
	return &RouterGroup{
		NotificationHandler:  notificationHandler,
		PreferenceHandler:    preferenceHandler,
		WebhookConfigHandler: webhookConfigHandler,
		SSEHandler:           sseHandler,
	}
}

// registerRoutes registers all HTTP routes for the Notification service.
func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	// Protected routes using JWT authentication.
	protected := r.Group("/")
	protected.Use(middlewares.Authentication(global.Config.JWT.PublicKey))

	// SSE (Server-Sent Events) - Stream is gin-coupled directly.
	protected.GET("/notifications/stream", rg.SSEHandler.Stream)

	// Notifications
	notifications := protected.Group("/notifications")
	{
		notifications.GET("", handler.Wrap(rg.NotificationHandler.Find))
		notifications.GET("/unread-count", handler.Wrap(rg.NotificationHandler.GetUnreadCount))
		notifications.PUT("/read-all", handler.Wrap(rg.NotificationHandler.MarkAllAsRead))
		notifications.PUT("/:id/read", handler.Wrap(rg.NotificationHandler.MarkAsRead))
	}

	// Webhooks
	webhooks := protected.Group("/webhooks")
	{
		webhooks.POST("", handler.Wrap(rg.WebhookConfigHandler.Create))
		webhooks.GET("", handler.Wrap(rg.WebhookConfigHandler.List))
		webhooks.GET("/:id", handler.Wrap(rg.WebhookConfigHandler.Get))
		webhooks.PUT("/:id", handler.Wrap(rg.WebhookConfigHandler.Update))
		webhooks.DELETE("/:id", handler.Wrap(rg.WebhookConfigHandler.Delete))
	}

	// User Preferences
	preferences := protected.Group("/preferences")
	{
		preferences.GET("", handler.Wrap(rg.PreferenceHandler.Get))
		preferences.PUT("", handler.Wrap(rg.PreferenceHandler.Update))
	}
}

// Ping health check endpoint.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Notification service running!",
	})
}

// NewEngine creates and configures the Gin engine.
func NewEngine(routerGroup *RouterGroup) *gin.Engine {
	if global.Config.Server.Mode != "release" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.Use(middlewares.RecoveryMiddleware)
	r.Use(middlewares.CORSMiddleware)

	r.GET("/ping", Ping)

	routerGroup.registerRoutes(r)

	return r
}
