package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DNSIntelResponse struct {
	gorm.Model
	DNSData datatypes.JSON `json:"a" gorm:"type:jsonb"` // Storing entire DNS response as JSON
	Host    string         `json:"host" gorm:"type:varchar(255)"`
}
