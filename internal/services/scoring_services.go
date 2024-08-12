package services

import "github.com/rkmangalp/bat-bet-go/internal/models"

type ScoringService struct{}

func NewScoringService() *ScoringService {
	return &ScoringService{}
}

func (ss *ScoringService) UpdateScores(match models.Match, result string) {
	if result == "WIN1_2" {
		match.Player1.Score++
		match.Player2.Score++
		match.Player3.Score--
		match.Player4.Score--
	} else {
		match.Player3.Score++
		match.Player4.Score++
		match.Player1.Score--
		match.Player2.Score--
	}
	match.Player1.Balance += float64(match.Player1.Score) * match.BetAmount
	match.Player2.Balance += float64(match.Player2.Score) * match.BetAmount
	match.Player3.Balance += float64(match.Player3.Score) * match.BetAmount
	match.Player4.Balance += float64(match.Player4.Score) * match.BetAmount
}
