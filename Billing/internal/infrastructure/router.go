package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-link/billing/global"
	driverHttp "go-link/billing/internal/adapters/driver/http"
	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"
	"go-link/common/pkg/permissions"
)

type RouterGroup struct {
	InvoiceHandler      driverHttp.InvoiceHandler
	PlanHandler         driverHttp.PlanHandler
	SubscriptionHandler driverHttp.SubscriptionHandler
}

func NewRouterGroup(
	invoiceHandler driverHttp.InvoiceHandler,
	planHandler driverHttp.PlanHandler,
	subscriptionHandler driverHttp.SubscriptionHandler,
) *RouterGroup {
	return &RouterGroup{
		InvoiceHandler:      invoiceHandler,
		PlanHandler:         planHandler,
		SubscriptionHandler: subscriptionHandler,
	}
}

func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	// Public Routes
	r.GET("/plans/active", handler.Wrap(rg.PlanHandler.FindActive))

	// Protected Routes
	protected := r.Group("/")
	protected.Use(middlewares.Authentication(global.Config.JWT.PublicKey))
	{
		// Admin Routes
		admin := protected.Group("", middlewares.RequireAdmin())
		{
			// Plans (Admin)
			admin.POST("/plans", handler.Wrap(rg.PlanHandler.Create))
			admin.PUT("/plans/:id", handler.Wrap(rg.PlanHandler.Update))
			admin.DELETE("/plans/:id", handler.Wrap(rg.PlanHandler.Delete))
			admin.GET("/plans", handler.Wrap(rg.PlanHandler.FindAll))
			admin.GET("/plans/:id", handler.Wrap(rg.PlanHandler.Get))

			// Invoices (Admin)
			admin.GET("/invoices/:id", handler.Wrap(rg.InvoiceHandler.Get))    // Get by ID
			admin.POST("/invoices", handler.Wrap(rg.InvoiceHandler.Create))    // Create
			admin.PUT("/invoices/:id", handler.Wrap(rg.InvoiceHandler.Update)) // Update
			admin.DELETE("/invoices/:id", handler.Wrap(rg.InvoiceHandler.Delete))

			// Subscriptions (Admin)
			admin.GET("/subscriptions/:id", handler.Wrap(rg.SubscriptionHandler.Get))
			admin.DELETE("/subscriptions/:id", handler.Wrap(rg.SubscriptionHandler.Delete))
		}

		// User Routes (Permission Based)
		// Invoices: FindMine (User Permission - Read)
		// "invoices thì viết thêm 2 api find theo user id ... dùng quyền"
		protected.GET("/invoices/mine", middlewares.RequirePermission(permissions.ResourceKeyInvoice, permissions.PermissionScopeRead), handler.Wrap(rg.InvoiceHandler.FindMine))

		// Subscriptions: Create/Update (User Permission)
		// "post, put là quyền"
		protected.POST("/subscriptions", middlewares.RequirePermission(permissions.ResourceKeySubscription, permissions.PermissionScopeCreate), handler.Wrap(rg.SubscriptionHandler.Create))
		protected.PUT("/subscriptions/:id", middlewares.RequirePermission(permissions.ResourceKeySubscription, permissions.PermissionScopeUpdate), handler.Wrap(rg.SubscriptionHandler.Update))
	}
}

// Ping health check endpoint.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Billing service running!",
	})
}

// NewEngine creates and configures the Gin engine.
func NewEngine(routerGroup *RouterGroup) *gin.Engine {
	if global.Config.Server.Mode != "release" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Middlewares
	r.Use(middlewares.CORSMiddleware)

	r.GET("/ping", Ping)

	// Register routes
	routerGroup.registerRoutes(r)

	return r
}
