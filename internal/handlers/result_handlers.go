package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkmangalp/bat-bet-go/internal/models"
	"github.com/rkmangalp/bat-bet-go/internal/services"
)

// ResultHandler handles HTTP requests related to match results
type ResultHandler struct {
	matchService   *services.MatchService
	scoringService *services.ScoringService
}

// NewResultHandler creates a new ResultHandler
func NewResultHandler(ms *services.MatchService, ss *services.ScoringService) *ResultHandler {
	return &ResultHandler{
		matchService:   ms,
		scoringService: ss,
	}
}

// UpdateResult handles updating the result of a match
func (rh *ResultHandler) UpdateResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID := vars["id"]

	var resultRequest models.Result
	if err := json.NewDecoder(r.Body).Decode(&resultRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := rh.matchService.UpdateResult(matchID, &resultRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rh.scoringService.UpdateScores(&resultRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
