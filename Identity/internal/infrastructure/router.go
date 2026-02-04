package infrastructure

import (
	"net/http"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"
	"go-link/common/pkg/permissions"

	"github.com/gin-gonic/gin"

	"go-link/identity/global"
	driverHttp "go-link/identity/internal/adapters/driver/http"
)

// RouterGroup contains all handlers.
type RouterGroup struct {
	TenantHandler              driverHttp.TenantHandler
	RoleHandler                driverHttp.RoleHandler
	PermissionHandler          driverHttp.PermissionHandler
	DomainHandler              driverHttp.DomainHandler
	ResourceHandler            driverHttp.ResourceHandler
	AttributeDefinitionHandler driverHttp.AttributeDefinitionHandler
	AuthenticationHandler      driverHttp.AuthenticationHandler
	UserHandler                driverHttp.UserHandler
}

// NewRouterGroup creates a new RouterGroup.
func NewRouterGroup(
	tenantHandler driverHttp.TenantHandler,
	roleHandler driverHttp.RoleHandler,
	permissionHandler driverHttp.PermissionHandler,
	domainHandler driverHttp.DomainHandler,
	resourceHandler driverHttp.ResourceHandler,
	attrDefHandler driverHttp.AttributeDefinitionHandler,
	authHandler driverHttp.AuthenticationHandler,
	userHandler driverHttp.UserHandler,
) *RouterGroup {
	return &RouterGroup{
		TenantHandler:              tenantHandler,
		RoleHandler:                roleHandler,
		PermissionHandler:          permissionHandler,
		DomainHandler:              domainHandler,
		ResourceHandler:            resourceHandler,
		AttributeDefinitionHandler: attrDefHandler,
		AuthenticationHandler:      authHandler,
		UserHandler:                userHandler,
	}
}

// registerRoutes registers all routes.
func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	r.POST("/users", handler.Wrap(rg.UserHandler.Create))

	// Public Routes
	public := r.Group("/auth")
	{
		public.POST("/register", handler.Wrap(rg.AuthenticationHandler.Register))
		public.POST("/login", handler.Wrap(rg.AuthenticationHandler.Login))
	}

	// Protected Routes
	protected := r.Group("/")
	protected.Use(middlewares.Authentication(global.Config.JWT.PublicKey))
	{
		// Auth
		auth := protected.Group("/auth")
		{
			auth.POST("/tenant-access", handler.Wrap(rg.AuthenticationHandler.AcquireToken))
			auth.POST("/change-password", handler.Wrap(rg.AuthenticationHandler.ChangePassword))
			auth.POST("/refresh", handler.Wrap(rg.AuthenticationHandler.RefreshToken))
		}

		// Tenants
		tenants := protected.Group("/tenants")
		{
			admin := tenants.Group("", middlewares.RequireAdmin())
			{
				admin.POST("", handler.Wrap(rg.TenantHandler.Create))
				admin.DELETE("/:id", handler.Wrap(rg.TenantHandler.Delete))
			}

			tenants.GET("/my", handler.Wrap(rg.TenantHandler.GetMyTenants))
			tenants.GET("/:id", middlewares.RequirePermission(permissions.ResourceKeyTenant, permissions.PermissionScopeRead), handler.Wrap(rg.TenantHandler.Get))
			tenants.PUT("/:id", middlewares.RequirePermission(permissions.ResourceKeyTenant, permissions.PermissionScopeUpdate), handler.Wrap(rg.TenantHandler.Update))
		}

		// Domains
		domains := protected.Group("/domains")
		{
			domains.POST("/find", middlewares.RequirePermission(permissions.ResourceKeyDomain, permissions.PermissionScopeRead), handler.Wrap(rg.DomainHandler.Find))
			domains.GET("/:id", middlewares.RequirePermission(permissions.ResourceKeyDomain, permissions.PermissionScopeRead), handler.Wrap(rg.DomainHandler.Get))
			domains.POST("", middlewares.RequirePermission(permissions.ResourceKeyDomain, permissions.PermissionScopeCreate), handler.Wrap(rg.DomainHandler.Create))
			domains.PUT("/:id", middlewares.RequirePermission(permissions.ResourceKeyDomain, permissions.PermissionScopeUpdate), handler.Wrap(rg.DomainHandler.Update))
			domains.DELETE("/:id", middlewares.RequirePermission(permissions.ResourceKeyDomain, permissions.PermissionScopeDelete), handler.Wrap(rg.DomainHandler.Delete))
		}

		// Roles
		roles := protected.Group("/roles", middlewares.RequireAdmin())
		{
			roles.POST("/find", handler.Wrap(rg.RoleHandler.Find))
			roles.GET("/:id", handler.Wrap(rg.RoleHandler.Get))
			roles.POST("", handler.Wrap(rg.RoleHandler.Create))
			roles.PUT("/:id", handler.Wrap(rg.RoleHandler.Update))
			roles.DELETE("/:id", handler.Wrap(rg.RoleHandler.Delete))
		}

		// Permissions
		permissions := protected.Group("/permissions", middlewares.RequireAdmin())
		{
			permissions.POST("/find", handler.Wrap(rg.PermissionHandler.Find))
			permissions.GET("/:id", handler.Wrap(rg.PermissionHandler.Get))
			permissions.POST("", handler.Wrap(rg.PermissionHandler.Create))
			permissions.PUT("/:id", handler.Wrap(rg.PermissionHandler.Update))
			permissions.DELETE("/:id", handler.Wrap(rg.PermissionHandler.Delete))
		}

		// Resources
		resources := protected.Group("/resources", middlewares.RequireAdmin())
		{
			resources.POST("/find", handler.Wrap(rg.ResourceHandler.Find))
			resources.GET("/:id", handler.Wrap(rg.ResourceHandler.Get))
			resources.POST("", handler.Wrap(rg.ResourceHandler.Create))
			resources.PUT("/:id", handler.Wrap(rg.ResourceHandler.Update))
			resources.DELETE("/:id", handler.Wrap(rg.ResourceHandler.Delete))
		}

		// AttributeDefinitions
		attrDefs := protected.Group("/attribute-definitions", middlewares.RequireAdmin())
		{
			attrDefs.POST("/find", handler.Wrap(rg.AttributeDefinitionHandler.Find))
			attrDefs.GET("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Get))
			attrDefs.POST("", handler.Wrap(rg.AttributeDefinitionHandler.Create))
			attrDefs.PUT("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Update))
			attrDefs.DELETE("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Delete))
		}

		// Users
		users := protected.Group("/users")
		{
			adminUsers := users.Group("", middlewares.RequireAdmin())
			{
				// adminUsers.POST("", handler.Wrap(rg.UserHandler.Create))
				adminUsers.DELETE("/:id", handler.Wrap(rg.UserHandler.Delete))
			}

			users.PUT("/profile", handler.Wrap(rg.UserHandler.UpdateProfile))
			users.GET("/profile", handler.Wrap(rg.UserHandler.GetProfile))
		}
	}
}

// Ping health check endpoint.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Identity service running!",
	})
}

// NewEngine creates and configures the Gin engine.
func NewEngine(routerGroup *RouterGroup) *gin.Engine {
	if global.Config.Server.Mode != "release" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Middlewares
	r.Use(middlewares.RecoveryMiddleware)
	r.Use(middlewares.CORSMiddleware)

	r.GET("/ping", Ping)

	// Register routes
	routerGroup.registerRoutes(r)

	return r
}
