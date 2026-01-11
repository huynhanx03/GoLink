package infrastructure

import (
	"net/http"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/common/http/middlewares"

	"github.com/gin-gonic/gin"

	"go-link/generation/global"
	driverHttp "go-link/generation/internal/adapters/driver/http"
)

// RouterGroup contains all routes
type RouterGroup struct {
	LinkHandler driverHttp.LinkHandler
}

// NewRouterGroup creates a new RouterGroup
func NewRouterGroup(
	linkHandler driverHttp.LinkHandler,
) *RouterGroup {
	return &RouterGroup{
		LinkHandler: linkHandler,
	}
}

// registerRoutes registers all routes
func (rg *RouterGroup) registerRoutes(r *gin.Engine) {
	links := r.Group("/links")
	{
		links.POST("", handler.Wrap(rg.LinkHandler.Create))
	}
}

// Ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "I'm running!",
	})
}

// NewEngine creates and configures the Gin engine
func NewEngine(routerGroup *RouterGroup) *gin.Engine {
	if global.Config.Server.Mode != "release" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// middlewares
	r.Use(middlewares.CORSMiddleware)

	r.GET("/ping", Ping)

	// Register routes
	routerGroup.registerRoutes(r)

	return r
}
