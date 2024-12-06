package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is an endpoint that returns a simple health check message
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
