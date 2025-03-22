package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dermot10/threat_intel_aggregator/internal/config"
	"github.com/dermot10/threat_intel_aggregator/internal/repositories"
	"github.com/dermot10/threat_intel_aggregator/internal/services"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	Repo *repositories.URLScanRepository
}

func NewURLHandler(repo *repositories.URLScanRepository) *UrlHandler {
	return &UrlHandler{Repo: repo}
}

func (h *UrlHandler) GetURLIntelHandler(c *gin.Context) {
	url := c.Query("url")

	if config.UrlScanKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is missing"})
		return
	}

	log.Println("API key found, calling service...")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	log.Println("Received URL:", url)

	response, err := services.PostUrlIntelService(url, config.UrlScanKey)
	if err != nil {
		log.Println("Error fetching URL Intel:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform URL scan"})
		return
	}

	apiResponse := *response

	if err := h.Repo.CreateURLIntel(&apiResponse); err != nil {
		log.Println("Error saving to DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save URL scan"})
		return
	}

	log.Printf("Received Data: %+v\n", response)
	c.JSON(http.StatusOK, response)
}

func (h *UrlHandler) GetStoredUrlIntelHandler(c *gin.Context) {
	url := c.Param("url")
	url = strings.TrimPrefix(url, "/")
	log.Println("Raw URL from request:", c.Param("url"))

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	data, err := h.Repo.GetURLIntel(url)
	if err != nil {
		log.Println("Error retrieving data from DB", err)
		return
	}

	log.Println("Retrieved data:", data)
	c.JSON(http.StatusOK, data)
}
