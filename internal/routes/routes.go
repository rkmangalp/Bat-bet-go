package routes

import (
	"github.com/gorilla/mux"
	"github.com/rkmangalp/bat-bet-go/internal/handlers"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

// NewRouter initializes and returns a new router with all the routes defined
func NewRouter(playerService *services.PlayerService, matchService *services.MatchService, scoringService *services.ScoringService) *mux.Router {
	// Initialize handlers
	playerHandler := handlers.NewPlayerHandler(playerService)
	matchHandler := handlers.NewMatchHandler(matchService, scoringService)
	resultHandler := handlers.NewResultHandler(matchService, scoringService)

	// Setup router
	router := mux.NewRouter()

	// Player routes
	router.HandleFunc("/players", playerHandler.AddPlayer).Methods("POST")
	router.HandleFunc("/players", playerHandler.GetPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")

	// Match routes
	router.HandleFunc("/matches", matchHandler.ScheduleMatch).Methods("POST")
	router.HandleFunc("/matches", matchHandler.GetMatches).Methods("GET")
	router.HandleFunc("/matches/{id}", matchHandler.GetMatchByID).Methods("GET")
	router.HandleFunc("/matches/{id}", matchHandler.UpdateMatch).Methods("PUT")
	router.HandleFunc("/matches/{id}", matchHandler.DeleteMatch).Methods("DELETE")

	// Result routes
	router.HandleFunc("/matches/{id}/result", resultHandler.UpdateResult).Methods("PUT")
	router.HandleFunc("/scoreboard", playerHandler.GetScoreboard).Methods("GET")

	return router
}
