package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dermot10/threat_intel_aggregator/internal/models"
)

const urlScanUrl = "https://urlscan.io/api/v1/scan/"

func PostUrlIntelService(url, apiKey string) (*models.UrlScanResponse, error) {
	requestURL := fmt.Sprintln(urlScanUrl, url)
	log.Println("Making request to:", requestURL)

	data := map[string]string{
		"url":        url,
		"visibility": "public",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlScanUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Println("Response Status Code:", resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	log.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	var result models.UrlScanResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
