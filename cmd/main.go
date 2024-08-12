package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkmangalp/bat-bet-go/internal/handlers"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

func main() {
	// Initialize services
	playerService := services.NewPlayerService()
	matchService := services.NewMatchService()
	scoringService := services.NewScoringService()

	// Initialize handlers
	playerHandler := handlers.NewPlayerHandler(playerService)
	matchHandler := handlers.NewMatchHandler(matchService, scoringService)
	resultHandler := handlers.NewResultHandler(matchService, scoringService)

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/players", playerHandler.AddPlayer).Methods("POST")
	router.HandleFunc("/matches", matchHandler.ScheduleMatch).Methods("POST")
	router.HandleFunc("/matches/{id}/result", resultHandler.UpdateResult).Methods("PUT")
	router.HandleFunc("/scoreboard", playerHandler.GetScoreboard).Methods("GET")

	// Start server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
