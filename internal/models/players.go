package models

import "github.com/google/uuid"

type Player struct {
	ID      uuid.UUID
	Name    string
	Score   int
	Balance float64
}
