package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AbuseIPKey string
var UrlScanKey string

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AbuseIPKey = os.Getenv("ABUSE_IPDB_KEY") // Load API key into global variable
	if AbuseIPKey == "" {
		log.Fatal("abuseIPDB_key key is missing in .env file")
	}
	UrlScanKey = os.Getenv("URL_SCAN_KEY")
	if UrlScanKey == "" {
		log.Fatal("url scan key is missing in .env file")
	}
}
