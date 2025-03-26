package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
	"gorm.io/gorm"
)

type DNSIntelRepository struct {
	DB *gorm.DB
}

func NewDNSIntelRepository(db *gorm.DB) *DNSIntelRepository {
	return &DNSIntelRepository{DB: db}
}

func (repo *DNSIntelRepository) CreateDNSIntel(apiResponse *models.DNSIntelResponse, domain string) error {
	if apiResponse == nil {
		return fmt.Errorf("received nil DNS intel response")
	}
	apiResponse.Host = domain

	dnsData, err := json.Marshal(apiResponse) // Marshal the DNS intel response into raw JSON
	if err != nil {
		return fmt.Errorf("failed to marshal DNS intel response: %w", err)
	}

	apiResponse.DNSData = json.RawMessage(dnsData) // Assign the marshaled data to the DNSData field in the database
	if err := repo.DB.Create(apiResponse).Error; err != nil {

		return fmt.Errorf("failed to insert DNS intel response: %w", err)
	}
	fmt.Printf("Successfully inserted DNS intel with ID: %d\n", apiResponse.ID)

	return nil
}
