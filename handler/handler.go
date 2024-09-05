package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlsortner/models"
	"urlsortner/services"
)

// Handler handles HTTP requests related to URL shortening
type Handler struct {
	Service services.Servicer
}

// New initializes the Handler
func New(service services.Servicer) *Handler {
	return &Handler{
		Service: service,
	}
}

// ShortenHandler handles the URL shortening requests
func (h *Handler) ShortenHandler(c *gin.Context) {
	var requestData models.RequestData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortURL := h.Service.ShortenURL(requestData.URL)

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// RedirectHandler handles redirection from short URLs to the original URLs
func (h *Handler) RedirectHandler(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, exists := h.Service.GetOriginalURL(shortURL)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
