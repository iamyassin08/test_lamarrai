package config

import "os"

// GetPort returns the port for the server to listen on
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	return port
}
