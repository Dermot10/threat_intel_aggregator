package repositories

import (
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
	// Ensure DNSData is not nil
	if len(apiResponse.DNSData) == 0 {
		return fmt.Errorf("apiResponse.DNSData is empty before saving")
	}
	fmt.Printf("Final JSON to be saved: %s\n", string(apiResponse.DNSData))

	if err := repo.DB.Create(apiResponse).Error; err != nil {
		return fmt.Errorf("failed to insert DNS intel response: %w", err)
	}

	fmt.Printf("Successfully inserted DNS intel with ID: %d\n", apiResponse.ID)
	return nil
}

func (repo *DNSIntelRepository) GetDNSIntel(host string) (*models.DNSIntelResponse, error) {
	var apiResponse models.DNSIntelResponse

	err := repo.DB.Where("host = ?", host).First(&apiResponse).Error
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
