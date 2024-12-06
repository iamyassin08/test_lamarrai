package router

import (
	"test_lamarria/be/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the API
func SetupRoutes(r *gin.Engine) {
	// Health check route
	r.GET("/health", handlers.HealthCheck)

	// Anomaly routes
	r.GET("/api/anomalies", handlers.GetAnomalies)
}
