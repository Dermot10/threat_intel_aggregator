package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB { //explicitly returns the connection object
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	fmt.Println("Using DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.ThreatIntel{},
		&models.AbuseIPDBResponse{},
		&models.Report{},
		&models.UrlScanResponse{},
		&models.UrlScanOptions{},
		&models.DNSIntelResponse{})

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	return db
}
