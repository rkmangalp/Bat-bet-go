package models

import "github.com/google/uuid"

type Match struct {
	ID        uuid.UUID
	Player1   Player
	Player2   Player
	Player3   Player
	Player4   Player
	Result    string // WIN1_2 or WIN3_4
	BetAmount float64
}
