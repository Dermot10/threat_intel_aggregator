package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
)

const dnsDumpsterURL = "https://api.dnsdumpster.com/domain/%s"

func GetDNSIntelService(domain, apikey string) (*models.DNSIntelResponse, error) {
	if apikey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	requestURL := fmt.Sprintf(dnsDumpsterURL, domain)
	log.Println("Making request to:", requestURL)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("X-API-Key", apikey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var data models.DNSIntelResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return &data, nil
}
