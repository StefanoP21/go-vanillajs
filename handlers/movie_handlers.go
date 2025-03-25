package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stefanop21/reelingit/logger"
	"github.com/stefanop21/reelingit/models"
)

type MovieHandler struct {
	logger *logger.Logger
}

// Utility functions
func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		}, {
			ID:          2,
			TMDB_ID:     182,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served top movies")
	}
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		}, {
			ID:          2,
			TMDB_ID:     182,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served random movies")
	}
}
