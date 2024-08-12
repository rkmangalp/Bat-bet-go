package models

import "github.com/google/uuid"

type Result struct {
	MatchID uuid.UUID
	Winner  string
	Loser   string
}
