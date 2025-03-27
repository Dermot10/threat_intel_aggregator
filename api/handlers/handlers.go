package handlers

import (
	"github.com/dermot10/threat_intel_aggregator/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Central routing function for main
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	IpRepo := repositories.NewThreatIntelRepository(db)
	IpHandler := NewIPIntelHandler(IpRepo)

	UrlRepo := repositories.NewUrlIntelRepository(db)
	UrlHandler := NewURLHandler(UrlRepo)

	DNSRepo := repositories.NewDNSIntelRepository(db)
	DNSHandler := NewDNSHandler(DNSRepo)

	router.GET("/", Homepage)
	router.GET("/intel/:ip", IpHandler.GetIPIntelHandler)
	router.GET("/stored-ipintel/:ip", IpHandler.GetStoredIPIntelHandler)
	router.POST("/url-scan", UrlHandler.CreateURLIntelHandler)
	router.GET("/stored-urlintel/*url", UrlHandler.GetStoredUrlIntelHandler)
	router.POST("/threat-intel", IpHandler.CreateIPIntelHandler)
	router.GET("/domain-intel/:domain", DNSHandler.CreateDNSIntelHandler)
	router.GET("/stored-domainintel/:host", DNSHandler.GetStoredDNSIntelHandler)
}
