package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dermot10/threat_intel_aggregator/internal/config"
	"github.com/dermot10/threat_intel_aggregator/internal/repositories"
	"github.com/dermot10/threat_intel_aggregator/internal/services"
	"github.com/gin-gonic/gin"
)

type DNSIntelHandler struct {
	Repo *repositories.DNSIntelRepository
}

func NewDNSHandler(repo *repositories.DNSIntelRepository) *DNSIntelHandler {
	return &DNSIntelHandler{Repo: repo}
}

func (h *DNSIntelHandler) CreateDNSIntelHandler(c *gin.Context) {
	domain := c.Param("domain")
	fmt.Println("Domain:", domain)

	if config.DNSDumpsterKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is missing"})
		return
	}

	response, err := services.GetDNSIntelService(domain, config.DNSDumpsterKey)
	if err != nil {
		log.Println("Error fetching DNS Intel:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch DNS intelligence"})
		return
	}

	apiResponse := *response

	if err := h.Repo.CreateDNSIntel(&apiResponse, domain); err != nil {
		log.Println("Error saving to DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save DNS intelligence"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DNS Intel record created", "id": apiResponse.ID})
}
