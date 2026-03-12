package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go-link/notification/global"
	"go-link/notification/internal/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server wraps HTTP
type Server struct {
	engine *gin.Engine
}

// NewServer
func NewServer(engine *gin.Engine) *Server {
	return &Server{
		engine: engine,
	}
}

// NewHTTPServer
func NewHTTPServer() *Server {
	c := di.GlobalContainer

	// Create router group with all handlers
	routerGroup := NewRouterGroup(
		c.NotificationContainer.Handler,
		c.PreferenceContainer.Handler,
		c.WebhookContainer.Handler,
		c.SSEContainer.Handler,
	)

	// Create Gin engine
	engine := NewEngine(routerGroup)

	// Create Server
	return NewServer(engine)
}

// Start HTTP server
func (s *Server) Start() (*http.Server, error) {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = global.Config.Server.Host
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(global.Config.Server.Port)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Addr:    address,
		Handler: s.engine,
	}

	go func() {
		global.LoggerZap.Info("HTTP Server starting",
			zap.String("address", address),
			zap.String("mode", global.Config.Server.Mode),
		)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LoggerZap.Fatal("Failed to start HTTP server", zap.Error(err))
		}
	}()

	return srv, nil
}
