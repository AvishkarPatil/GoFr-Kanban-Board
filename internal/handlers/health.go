package handlers

import (
	"gofr.dev/pkg/gofr"
)

func HealthCheck(c *gofr.Context) (interface{}, error) {
	// Test database connection
	_, err := c.SQL.Query("SELECT 1")
	if err != nil {
		return map[string]interface{}{
			"status":   "unhealthy",
			"database": "disconnected",
			"error":    err.Error(),
		}, nil
	}

	return map[string]interface{}{
		"status":   "healthy",
		"database": "connected",
		"service":  "kanban-board",
	}, nil
}