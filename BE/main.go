package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Anomaly represents a window anomaly
type Anomaly struct {
	ID       int    `json:"id"`
	Window   string `json:"window"`
	Severity string `json:"severity"`
}

// Fetch anomaly data from a mock JSON file
func getAnomalies(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("anomalies.json")
	if err != nil {
		http.Error(w, "Could not open data file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var anomalies []Anomaly
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&anomalies)
	if err != nil {
		http.Error(w, "Could not decode data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anomalies)
}

func main() {
	http.HandleFunc("/api/anomalies", getAnomalies)

	// Serve static files (frontend)
	http.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
