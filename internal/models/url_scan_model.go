package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UrlScanOptions struct {
	UserAgent string `json:"useragent"`
}

type UrlScanResponse struct {
	gorm.Model
	Message    string         `json:"message"`
	UUID       string         `json:"uuid"`
	Result     string         `json:"result"`
	API        string         `json:"api"`
	Visibility string         `json:"visibility"`
	Options    datatypes.JSON `json:"options"`
	Url        string         `json:"url"`
	Country    string         `json:"country"`
}
