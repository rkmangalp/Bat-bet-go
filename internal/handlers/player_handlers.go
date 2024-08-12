package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rkmangalp/bat-bet-go/internal/services"
)

type PlayerHandler struct {
	playerService *services.PlayerService
}

func NewPlayerHandler(ps *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: ps}
}

func (ph *PlayerHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	var name struct{ Name string }
	json.NewDecoder(r.Body).Decode(&name)
	player := ph.playerService.AddPlayer(name.Name)
	json.NewEncoder(w).Encode(player)
}
