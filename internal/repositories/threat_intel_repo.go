package repositories

import (
	"fmt"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ThreatIntelRepository struct {
	DB *gorm.DB
}

func NewThreatIntelRepository(db *gorm.DB) *ThreatIntelRepository {
	return &ThreatIntelRepository{DB: db}
}

func (repo *ThreatIntelRepository) CreateThreatIntel(intel *models.ThreatIntel) error {
	return repo.DB.Create(intel).Error
}

// api response of type model #pointer to struct
func (repo *ThreatIntelRepository) CreateIPIntel(apiResponse *models.AbuseIPDBAPIResponse) error {
	ipIntel := apiResponse.Data
	//insert into db, deference to intel pointer
	if err := repo.DB.Create(&ipIntel).Error; err != nil {
		fmt.Println("DB insert error:", err)
		return err
	}

	fmt.Println("Inserted IP intel ID:", ipIntel.ID)

	//creates, if error assign to err and check if not nil
	for i := range ipIntel.Reports {
		ipIntel.Reports[i].AbuseIPDBResponseID = ipIntel.ID

		if err := repo.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}}, // If "id" is duplicated
			DoNothing: true,                          // Ignore the insert instead of failing
		}).Create(&ipIntel.Reports[i]).Error; err != nil {
			fmt.Println("Skipping duplicate report ID:", ipIntel.Reports[i].ID)
		}
	}

	//indicates success
	return nil
}

func (repo *ThreatIntelRepository) GetIPIntel(ip string) (*models.AbuseIPDBResponse, error) {
	var apiResponse models.AbuseIPDBResponse //empty struct, to hold db result

	//populates struct with response of db query, if found
	err := repo.DB.Where("ip_address = ?", ip).First(&apiResponse).Error
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
