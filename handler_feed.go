package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "github.com/mide7/go/rss_aggregator/db/sqlc"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing request: ", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		Title:  params.Title,
		Url:    params.Url,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error creating user: ", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

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

	log.Printf("limit: %d, page: %d, offset: %d", limit, page, offset)

	feeds, err := apiCfg.DB.ListFeeds(r.Context(), db.ListFeedsParams{
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
		"data":  databaseFeedsToFeeds(feeds),
	})
}
