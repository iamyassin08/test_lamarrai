package storage

import (
	"encoding/json"
	"os"
	"test_lamarria/be/models"
)

// FetchAnomalies retrieves the list of anomalies from the mock data
func FetchAnomalies() ([]models.Anomaly, error) {
	var anomalies []models.Anomaly

	// Open the mock data file
	file, err := os.Open("storage/mock_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON data
	err = json.NewDecoder(file).Decode(&anomalies)
	if err != nil {
		return nil, err
	}

	return anomalies, nil
}
