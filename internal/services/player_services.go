package services

import (
	"github.com/rkmangalp/bat-bet-go/internal/models"

	"github.com/google/uuid"
)

type PlayerService struct {
	players []models.Player
}

func NewPlayerService() *PlayerService {
	return &PlayerService{
		players: []models.Player{},
	}
}

func (ps *PlayerService) AddPlayer(name string) models.Player {
	player := models.Player{
		ID:      uuid.New(),
		Name:    name,
		Score:   0,
		Balance: 0.0,
	}
	ps.players = append(ps.players, player)
	return player
}

func (ps *PlayerService) GetPlayers() []models.Player {
	return ps.players
}
