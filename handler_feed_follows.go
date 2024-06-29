package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/maxkopitz/rss-feed-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get follow feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedsFollows))
}

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode paramaters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})

}
