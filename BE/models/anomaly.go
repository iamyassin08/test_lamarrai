package models

// Anomaly defines the structure of an anomaly record
type Anomaly struct {
	ID       int    `json:"id"`
	Window   string `json:"window"`
	Severity string `json:"severity"`
}
