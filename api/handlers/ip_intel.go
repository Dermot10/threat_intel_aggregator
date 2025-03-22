package handlers

import (
	"log"
	"net/http"

	"github.com/dermot10/threat_intel_aggregator/internal/config"
	"github.com/dermot10/threat_intel_aggregator/internal/models"
	"github.com/dermot10/threat_intel_aggregator/internal/repositories"
	"github.com/dermot10/threat_intel_aggregator/internal/services"
	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is the threat intel homepage",
	})
}

type IntelHandler struct {
	Repo *repositories.ThreatIntelRepository
}

func NewIntelHandler(repo *repositories.ThreatIntelRepository) *IntelHandler {
	return &IntelHandler{Repo: repo}
}

func (h *IntelHandler) CreateIntelHandler(c *gin.Context) {
	var intel models.ThreatIntel

	// bind json input data from the handler to intel struct which is created from the model
	if err := c.ShouldBindJSON(&intel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	//crud func here
	if err := h.Repo.CreateThreatIntel(&intel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store threat intel"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Threat intel successfully created", "intel": intel})
}

func (h *IntelHandler) GetIPIntelHandler(c *gin.Context) {
	ip := c.Param("ip")

	if config.AbuseIPKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is missing"})
		return
	}

	response, err := services.GetIPIntelService(ip, config.AbuseIPKey)
	if err != nil {
		log.Println("Error fetching Ip Intel:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch IP intelligence"})
		return
	}

	apiResponse := models.AbuseIPDBAPIResponse{Data: *response}

	if err := h.Repo.CreateIPIntel(&apiResponse); err != nil {
		log.Println("Error saving to DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save IP intelligence"})
		return
	}

	log.Printf("Received Data: %+v\n", response)
	c.JSON(http.StatusOK, response)
}

func (h *IntelHandler) GetStoredIPIntelHandler(c *gin.Context) {
	ip := c.Param("ip")
	log.Println("Handler called for IP:", ip)

	data, err := h.Repo.GetIPIntel(ip)
	if err != nil {
		log.Println("Error retrieving data from DB", err)
		return
	}
	log.Println("Retrieved data:", data)
	c.JSON(http.StatusOK, data)
}
