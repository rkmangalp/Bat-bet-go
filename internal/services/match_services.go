package services

import (
	"errors"
	"sync"

	"github.com/rkmangalp/bat-bet-go/internal/models"
)

// MatchService provides operations related to matches
type MatchService struct {
	matches map[string]*models.Match
	mu      sync.Mutex
}

// NewMatchService creates a new MatchService
func NewMatchService() *MatchService {
	return &MatchService{
		matches: make(map[string]*models.Match),
	}
}

// ScheduleMatch schedules a new match
func (ms *MatchService) ScheduleMatch(match *models.Match) *models.Match {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	match.ID = generateID()
	ms.matches[match.ID] = match
	return match
}

// GetMatches retrieves all matches
func (ms *MatchService) GetMatches() []*models.Match {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	matches := []*models.Match{}
	for _, match := range ms.matches {
		matches = append(matches, match)
	}
	return matches
}

// GetMatchByID retrieves a match by its ID
func (ms *MatchService) GetMatchByID(id string) (*models.Match, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	match, exists := ms.matches[id]
	if !exists {
		return nil, errors.New("match not found")
	}
	return match, nil
}

// UpdateMatch updates a match's details
func (ms *MatchService) UpdateMatch(id string, matchUpdate *models.Match) (*models.Match, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	match, exists := ms.matches[id]
	if !exists {
		return nil, errors.New("match not found")
	}

	match.Team1 = matchUpdate.Team1
	match.Team2 = matchUpdate.Team2
	match.Result = matchUpdate.Result

	return match, nil
}

// DeleteMatch deletes a match by its ID
func (ms *MatchService) DeleteMatch(id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	_, exists := ms.matches[id]
	if !exists {
		return errors.New("match not found")
	}

	delete(ms.matches, id)
	return nil
}

// UpdateResult updates the result of a match and returns any error encountered
func (ms *MatchService) UpdateResult(id string, result *models.Result) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	match, exists := ms.matches[id]
	if !exists {
		return errors.New("match not found")
	}

	match.Result = result
	return nil
}
