package main

import (
	"log"
	"net/http"

	"github.com/rkmangalp/bat-bet-go/internal/config"
	"github.com/rkmangalp/bat-bet-go/internal/routes"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

func main() {
	// Load configuration
	config.LoadConfig("config/config.json")

	// Initialize services
	playerService := services.NewPlayerService()
	matchService := services.NewMatchService()
	scoringService := services.NewScoringService()

	// Initialize router
	router := routes.NewRouter(playerService, matchService, scoringService)

	// Start server
	log.Printf("Starting server on port %s", config.AppConfig.Port)
	err := http.ListenAndServe(":"+config.AppConfig.Port, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
