package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkmangalp/bat-bet-go/internal/models"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

// MatchHandler handles HTTP requests related to matches
type MatchHandler struct {
	matchService   *services.MatchService
	scoringService *services.ScoringService
}

// NewMatchHandler creates a new MatchHandler
func NewMatchHandler(ms *services.MatchService, ss *services.ScoringService) *MatchHandler {
	return &MatchHandler{
		matchService:   ms,
		scoringService: ss,
	}
}

// ScheduleMatch handles scheduling a new match
func (mh *MatchHandler) ScheduleMatch(w http.ResponseWriter, r *http.Request) {
	var matchRequest models.Match
	if err := json.NewDecoder(r.Body).Decode(&matchRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	match := mh.matchService.ScheduleMatch(&matchRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// GetMatches handles retrieving all matches
func (mh *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches := mh.matchService.GetMatches()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

// GetMatchByID handles retrieving a match by ID
func (mh *MatchHandler) GetMatchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID := vars["id"]

	match, err := mh.matchService.GetMatchByID(matchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// UpdateMatch handles updating match details
func (mh *MatchHandler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID := vars["id"]

	var matchUpdateRequest models.Match
	if err := json.NewDecoder(r.Body).Decode(&matchUpdateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	match, err := mh.matchService.UpdateMatch(matchID, &matchUpdateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// DeleteMatch handles deleting a match by ID
func (mh *MatchHandler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID := vars["id"]

	err := mh.matchService.DeleteMatch(matchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
