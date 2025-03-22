package handlers

import (
	"github.com/dermot10/threat_intel_aggregator/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Central routing function for main
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	IpRepo := repositories.NewThreatIntelRepository(db)
	IpHandler := NewIntelHandler(IpRepo)

	UrlRepo := repositories.NewUrlIntelRepository(db)
	UrlHandler := NewURLHandler(UrlRepo)

	router.GET("/", Homepage)
	router.GET("/intel/:ip", IpHandler.GetIPIntelHandler)
	router.GET("/stored-intel/:ip", IpHandler.GetStoredIPIntelHandler)
	router.POST("/url-scan", UrlHandler.GetURLIntelHandler)
	router.GET("/stored-url/*url", UrlHandler.GetStoredUrlIntelHandler)
	router.POST("/threat-intel", IpHandler.CreateIntelHandler)

}
