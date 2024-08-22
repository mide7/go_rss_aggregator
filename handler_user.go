package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/mide7/go_rss_aggregator/db/sqlc"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing request: ", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), params.Name)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error creating user: ", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := apiCfg.DB.ListUsers(r.Context(), db.ListUsersParams{
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting users: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUsersToUsers(users))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUserPosts(w http.ResponseWriter, r *http.Request, user db.User) {
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

	posts, err := apiCfg.DB.ListPostsForUser(r.Context(), db.ListPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: offset,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting posts: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]any{
		"data":  databasePostsToPosts(posts),
		"page":  page,
		"limit": limit,
	})
}
