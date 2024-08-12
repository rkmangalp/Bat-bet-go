package services

import (
	"github.com/rkmangalp/bat-bet-go/internal/models"

	"github.com/google/uuid"
)

type MatchService struct {
	matches []models.Match
}

func NewMatchService() *MatchService {
	return &MatchService{
		matches: []models.Match{},
	}
}

func (ms *MatchService) ScheduleMatch(players []models.Player, betAmount float64) models.Match {
	match := models.Match{
		ID:        uuid.New(),
		Player1:   players[0],
		Player2:   players[1],
		Player3:   players[2],
		Player4:   players[3],
		BetAmount: betAmount,
	}
	ms.matches = append(ms.matches, match)
	return match
}

func (ms *MatchService) GetMatches() []models.Match {
	return ms.matches
}
