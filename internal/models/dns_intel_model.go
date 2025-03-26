package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type DNSIntelResponse struct {
	gorm.Model
	DNSData json.RawMessage `json:"dns_data" gorm:"type:jsonb"` // Storing entire DNS response as JSON
	Host    string          `json:"host" gorm:"type:varchar(255)"`
}
