package main

import (
	"github.com/AvishkarPatil/GoFr-Kanban-Board/internal/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Health check endpoint
	app.GET("/health", handlers.HealthCheck)

	// Start server
	app.Run()
}