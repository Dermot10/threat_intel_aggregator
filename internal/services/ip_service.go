package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
)

const abuseIPDBURL = "https://api.abuseipdb.com/api/v2/check"

func GetIPIntelService(ip, apiKey string) (*models.AbuseIPDBResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", abuseIPDBURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %w", err)
	}

	query := req.URL.Query()
	query.Add("ipAddress", ip)
	query.Add("maxAgeInDays", "90")
	query.Add("verbose", "true")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Key", apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	var apiResponse models.AbuseIPDBAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &apiResponse.Data, nil
}
