package models

import (
	"time"

	"github.com/lib/pq"

	"gorm.io/gorm"
)

type AbuseIPDBAPIResponse struct {
	Data AbuseIPDBResponse `json:"data"`
}

type Report struct {
	ID                  uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	AbuseIPDBResponseID uint          `gorm:"not null;index" json:"abuseIpDbResponseId"`
	ReportedAt          time.Time     `json:"reportedAt"`
	Comment             string        `json:"comment"`
	Categories          pq.Int32Array `gorm:"type:integer[]" json:"categories"`
	ReporterID          int           `json:"reporterId"`
	ReporterCountryCode string        `json:"reporterCountryCode"`
	ReporterCountryName string        `json:"reporterCountryName"`
}

// Struct to represent the response from the IP intelligence API
type AbuseIPDBResponse struct {
	gorm.Model
	IPAddress            string         `json:"ipAddress"`
	IsPublic             bool           `json:"isPublic"`
	IpVersion            int            `json:"ipVersion"`
	IsWhiteListed        bool           `json:"isWhitelisted"`
	AbuseConfidenceScore int            `json:"abuseConfidenceScore"`
	CountryCode          string         `json:"countryCode"`
	CountryName          string         `json:"countryName"`
	UsageType            string         `json:"usageType"`
	ISP                  string         `json:"isp"`
	Domain               string         `json:"domain"`
	Hostnames            pq.StringArray `gorm:"type:text[]" json:"hostnames"`
	IsTor                bool           `json:"isTor"`
	TotalReports         int            `json:"totalReports"`
	LastReportedAt       *time.Time     `json:"lastReportedAt"`
	Reports              []Report       `gorm:"foreignKey:AbuseIPDBResponseID" json:"reports"`
}

type ThreatIntel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Indicator   string `json:"indicator"`
	Type        string `json:"type"`
	IntelSource string `json:"intel_source"`
	Description string `json:"description"`
}
