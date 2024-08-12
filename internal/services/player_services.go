package services

import (
	"errors"
	"sync"

	"github.com/rkmangalp/bat-bet-go/internal/models"
)

// PlayerService provides operations related to players
type PlayerService struct {
	players map[string]*models.Player
	mu      sync.Mutex
}

// NewPlayerService creates a new PlayerService
func NewPlayerService() *PlayerService {
	return &PlayerService{
		players: make(map[string]*models.Player),
	}
}

// AddPlayer adds a new player
func (ps *PlayerService) AddPlayer(name string) *models.Player {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	player := &models.Player{
		ID:      generateID(),
		Name:    name,
		Score:   0,
		Balance: 0,
	}

	ps.players[player.ID] = player
	return player
}

// GetPlayers retrieves all players
func (ps *PlayerService) GetPlayers() []*models.Player {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	players := []*models.Player{}
	for _, player := range ps.players {
		players = append(players, player)
	}
	return players
}

// GetPlayerByID retrieves a player by their ID
func (ps *PlayerService) GetPlayerByID(id string) (*models.Player, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	player, exists := ps.players[id]
	if !exists {
		return nil, errors.New("player not found")
	}
	return player, nil
}

// UpdatePlayer updates a player's information
func (ps *PlayerService) UpdatePlayer(id string, playerUpdate *models.Player) (*models.Player, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	player, exists := ps.players[id]
	if !exists {
		return nil, errors.New("player not found")
	}

	player.Name = playerUpdate.Name
	player.Score = playerUpdate.Score
	player.Balance = playerUpdate.Balance

	return player, nil
}

// DeletePlayer deletes a player by their ID
func (ps *PlayerService) DeletePlayer(id string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	_, exists := ps.players[id]
	if !exists {
		return errors.New("player not found")
	}

	delete(ps.players, id)
	return nil
}

// GetScoreboard retrieves the current scoreboard
func (ps *PlayerService) GetScoreboard() []*models.Player {
	return ps.GetPlayers()
}
