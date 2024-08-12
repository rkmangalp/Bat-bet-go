package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkmangalp/bat-bet-go/internal/models"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

// PlayerHandler handles HTTP requests related to players
type PlayerHandler struct {
	playerService *services.PlayerService
}

// NewPlayerHandler creates a new PlayerHandler
func NewPlayerHandler(ps *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: ps}
}

// AddPlayer handles adding a new player
func (ph *PlayerHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	var playerRequest struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&playerRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player := ph.playerService.AddPlayer(playerRequest.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// GetPlayers handles retrieving all players
func (ph *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := ph.playerService.GetPlayers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

// GetPlayerByID handles retrieving a player by ID
func (ph *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["id"]

	player, err := ph.playerService.GetPlayerByID(playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// UpdatePlayer handles updating a player's information
func (ph *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["id"]

	var playerUpdateRequest models.Player
	if err := json.NewDecoder(r.Body).Decode(&playerUpdateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := ph.playerService.UpdatePlayer(playerID, &playerUpdateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// DeletePlayer handles deleting a player by ID
func (ph *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["id"]

	err := ph.playerService.DeletePlayer(playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetScoreboard handles retrieving the scoreboard
func (ph *PlayerHandler) GetScoreboard(w http.ResponseWriter, r *http.Request) {
	scoreboard := ph.playerService.GetScoreboard()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scoreboard)
}
