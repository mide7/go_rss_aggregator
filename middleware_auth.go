package main

import (
	"fmt"
	"net/http"

	db "github.com/mide7/go_rss_aggregator/db/sqlc"
	"github.com/mide7/go_rss_aggregator/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprint("Error getting user: ", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting user: ", err))
			return
		}

		handler(w, r, user)
	}
}
