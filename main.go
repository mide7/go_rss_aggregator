package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	db "github.com/mide7/go/rss_aggregator/db/sqlc"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if port == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	conn, err := sql.Open("postgres", databaseURL)

	db := db.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/health-check", handlerReadiness)
	v1Router.Get("/users", apiCfg.handleGetUsers)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/me", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Get("/feeds", apiCfg.handleGetFeeds)
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))

	v1Router.Get("/feed-follows", apiCfg.handleGetFeedFollows)
	v1Router.Get("/feed-follows/me", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollowsByUser))
	v1Router.Post("/feed-follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Delete("/feed-follows/{id}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetUserPosts))

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Println("\n ========================= \n Listening on port:", port, "\n =========================")
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
