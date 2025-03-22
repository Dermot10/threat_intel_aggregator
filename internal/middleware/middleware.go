package config

import (
	"net/http"

	"github.com/dermot10/threat_intel_aggregator/internal/config"
	"github.com/gin-gonic/gin"
)

// run before handlers/routes
func RequireAbuseIPDBKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AbuseIPKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"errror": "AbuseIPDB API key is missing"})
			c.Abort()
			return
		}
		c.Next()
	}
}
