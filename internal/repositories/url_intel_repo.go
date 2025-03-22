package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type URLScanRepository struct {
	DB *gorm.DB
}

func NewUrlIntelRepository(db *gorm.DB) *URLScanRepository {
	return &URLScanRepository{DB: db}
}

func (repo *URLScanRepository) CreateURLIntel(apiResponse *models.UrlScanResponse) error {

	optionsJSON, err := json.Marshal(apiResponse.Options)
	if err != nil {
		return fmt.Errorf("failed to marshal options: %w", err)
	}

	urlIntel := *apiResponse
	urlIntel.Options = datatypes.JSON(optionsJSON)

	if err := repo.DB.Create(&urlIntel).Error; err != nil {
		fmt.Println("DB insert error:", err)
		return err
	}
	fmt.Println("Inserted URL intel:", urlIntel.ID)

	//indicates success
	return nil

}

func (repo *URLScanRepository) GetURLIntel(url string) (*models.UrlScanResponse, error) {
	var apiResponse models.UrlScanResponse
	err := repo.DB.Where("url = ?", url).First(&apiResponse).Error
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
