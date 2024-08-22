package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	db "github.com/mide7/go_rss_aggregator/db/sqlc"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing request: ", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error following feed : ", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request) {

	// get query params
	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {
		limitParam = "10"
	}
	limit, err := strconv.ParseInt(limitParam, 10, 64)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing limit query param")
		return
	}

	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.ParseInt(pageParam, 10, 64)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing page query param")
		return
	}

	offset := (int32(page) - 1) * int32(limit)

	feedFollows, err := apiCfg.DB.ListFeedFollows(r.Context(), db.ListFeedFollowsParams{
		Limit:  int32(limit),
		Offset: offset,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting feeds: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"page":  page,
		"limit": limit,
		"data":  databaseListFeedFollowsToFeedFollows(feedFollows),
	})
}

func (apiCfg *apiConfig) handleGetFeedFollowsByUser(w http.ResponseWriter, r *http.Request, user db.User) {
	// get query params
	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {
		limitParam = "10"
	}
	limit, err := strconv.ParseInt(limitParam, 10, 64)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing limit query param")
		return
	}

	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.ParseInt(pageParam, 10, 64)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing page query param")
		return
	}

	offset := (int32(page) - 1) * int32(limit)

	feedFollows, err := apiCfg.DB.ListFeedFollowsByUserID(r.Context(), db.ListFeedFollowsByUserIDParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: offset,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting feeds: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"page":  page,
		"limit": limit,
		"data":  databaseFeedFollowsToFeedFollows(feedFollows),
	})
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	idString := chi.URLParam(r, "id")

	ID, err := uuid.Parse(idString)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing ID: ", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error deleting feed follow: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
