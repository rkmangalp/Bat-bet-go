package services

import (
	"errors"
	"sync"

	"github.com/rkmangalp/bat-bet-go/internal/models"
)

// ScoringService provides operations related to scoring and balances
type ScoringService struct {
	playerService *PlayerService
	mu            sync.Mutex
}

// NewScoringService creates a new ScoringService
func NewScoringService(ps *PlayerService) *ScoringService {
	return &ScoringService{
		playerService: ps,
	}
}

// UpdateScores updates the scores and balances of players based on match results
func (ss *ScoringService) UpdateScores(result *models.Result) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	for _, winnerID := range result.Winners {
		player, err := ss.playerService.GetPlayerByID(winnerID)
		if err != nil {
			return errors.New("error updating winner's score: " + err.Error())
		}
		player.Score += 1
		player.Balance += result.BetAmount
	}

	for _, loserID := range result.Losers {
		player, err := ss.playerService.GetPlayerByID(loserID)
		if err != nil {
			return errors.New("error updating loser's score: " + err.Error())
		}
		player.Score -= 1
		player.Balance -= result.BetAmount
	}

	return nil
}
