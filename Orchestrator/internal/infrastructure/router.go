package infrastructure

import (
	"net/http"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"

	"github.com/gin-gonic/gin"

	"go-link/orchestrator/global"
	driverHttp "go-link/orchestrator/internal/adapters/driver/http"
)

// RouterGroup contains all handlers.
type RouterGroup struct {
	OrchestratorHandler driverHttp.OrchestratorHandler
}

// NewRouterGroup creates a new RouterGroup.
func NewRouterGroup(
	orchestratorHandler driverHttp.OrchestratorHandler,
) *RouterGroup {
	return &RouterGroup{
		OrchestratorHandler: orchestratorHandler,
	}
}

// registerRoutes registers all routes.
func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	// Public Routes
	public := r.Group("/auth")
	{
		public.POST("/register", handler.Wrap(rg.OrchestratorHandler.Register))
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
