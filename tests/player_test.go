package tests

import (
	"testing"

	"github.com/rkmangalp/bat-bet-go/internal/services"
)

func TestAddPlayer(t *testing.T) {
	playerService := services.NewPlayerService()
	player := playerService.AddPlayer("Player 1")
	if player.Name != "Player 1" {
		t.Errorf("Expected 'Player 1', got %s", player.Name)
	}
}
