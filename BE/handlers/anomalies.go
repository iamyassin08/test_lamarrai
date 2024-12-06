package handlers

import (
	"net/http"
	"test_lamarria/be/storage"

	"github.com/gin-gonic/gin"
)

// GetAnomalies retrieves the list of anomalies from the mock data
func GetAnomalies(c *gin.Context) {
	anomalies, err := storage.FetchAnomalies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch anomalies"})
		return
	}

	c.JSON(http.StatusOK, anomalies)
}
