package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-link/common/pkg/common/http/handler"
	"go-link/redirection/internal/ports"
)

type LinkHandler interface {
	Redirect(c *gin.Context)
}

type linkHandler struct {
	handler.BaseHandler
	linkService ports.LinkService
}

func NewLinkHandler(linkService ports.LinkService) LinkHandler {
	return &linkHandler{
		linkService: linkService,
	}
}

// Redirect handles the redirection to original URL
func (h *linkHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid short code"})
		return
	}

	url, err := h.linkService.GetOriginalURL(c.Request.Context(), shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
