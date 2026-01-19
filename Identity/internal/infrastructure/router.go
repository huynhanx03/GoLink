package infrastructure

import (
	"net/http"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"

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
}

// NewRouterGroup creates a new RouterGroup.
func NewRouterGroup(
	tenantHandler driverHttp.TenantHandler,
	roleHandler driverHttp.RoleHandler,
	permissionHandler driverHttp.PermissionHandler,
	domainHandler driverHttp.DomainHandler,
	resourceHandler driverHttp.ResourceHandler,
	attrDefHandler driverHttp.AttributeDefinitionHandler,
) *RouterGroup {
	return &RouterGroup{
		TenantHandler:              tenantHandler,
		RoleHandler:                roleHandler,
		PermissionHandler:          permissionHandler,
		DomainHandler:              domainHandler,
		ResourceHandler:            resourceHandler,
		AttributeDefinitionHandler: attrDefHandler,
	}
}

// registerRoutes registers all routes.
func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	// Tenant routes
	tenants := r.Group("/tenants")
	{
		tenants.GET("/:id", handler.Wrap(rg.TenantHandler.Get))
		tenants.POST("", handler.Wrap(rg.TenantHandler.Create))
		tenants.PUT("/:id", handler.Wrap(rg.TenantHandler.Update))
		tenants.DELETE("/:id", handler.Wrap(rg.TenantHandler.Delete))
	}

	// Role routes
	roles := r.Group("/roles")
	{
		roles.POST("/find", handler.Wrap(rg.RoleHandler.Find))
		roles.GET("/:id", handler.Wrap(rg.RoleHandler.Get))
		roles.POST("", handler.Wrap(rg.RoleHandler.Create))
		roles.PUT("/:id", handler.Wrap(rg.RoleHandler.Update))
		roles.DELETE("/:id", handler.Wrap(rg.RoleHandler.Delete))
	}

	// Permission routes
	permissions := r.Group("/permissions")
	{
		permissions.POST("/find", handler.Wrap(rg.PermissionHandler.Find))
		permissions.GET("/:id", handler.Wrap(rg.PermissionHandler.Get))
		permissions.POST("", handler.Wrap(rg.PermissionHandler.Create))
		permissions.PUT("/:id", handler.Wrap(rg.PermissionHandler.Update))
		permissions.DELETE("/:id", handler.Wrap(rg.PermissionHandler.Delete))
	}

	// Domain routes
	domains := r.Group("/domains")
	{
		domains.POST("/find", handler.Wrap(rg.DomainHandler.Find))
		domains.GET("/:id", handler.Wrap(rg.DomainHandler.Get))
		domains.POST("", handler.Wrap(rg.DomainHandler.Create))
		domains.PUT("/:id", handler.Wrap(rg.DomainHandler.Update))
		domains.DELETE("/:id", handler.Wrap(rg.DomainHandler.Delete))
	}

	// Resource routes
	resources := r.Group("/resources")
	{
		resources.POST("/find", handler.Wrap(rg.ResourceHandler.Find))
		resources.GET("/:id", handler.Wrap(rg.ResourceHandler.Get))
		resources.POST("", handler.Wrap(rg.ResourceHandler.Create))
		resources.PUT("/:id", handler.Wrap(rg.ResourceHandler.Update))
		resources.DELETE("/:id", handler.Wrap(rg.ResourceHandler.Delete))
	}

	// AttributeDefinition routes
	attrDefs := r.Group("/attribute-definitions")
	{
		attrDefs.POST("/find", handler.Wrap(rg.AttributeDefinitionHandler.Find))
		attrDefs.GET("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Get))
		attrDefs.POST("", handler.Wrap(rg.AttributeDefinitionHandler.Create))
		attrDefs.PUT("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Update))
		attrDefs.DELETE("/:id", handler.Wrap(rg.AttributeDefinitionHandler.Delete))
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
	r.Use(middlewares.CORSMiddleware)

	r.GET("/ping", Ping)

	// Register routes
	routerGroup.registerRoutes(r)

	return r
}
