package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lovesupergames/RSSAgregator/internal/api"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		return
	}
	port := os.Getenv("PORT")
	conStr := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	dbQueries := database.New(db)
	config := api.ApiConfig{
		dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /app/v1/healthz", common.HealthCheck)
	mux.HandleFunc("GET /app/v1/err", common.ErrCheck)
	mux.HandleFunc("POST /app/v1/users", config.HandleCreateUser)
	mux.HandleFunc("GET /app/v1/users", config.MiddlewareAuthUser(config.HandleGetUsers))
	mux.HandleFunc("POST /app/v1/feeds", config.MiddlewareAuthUser(config.HandlePostFeed))
	mux.HandleFunc("GET /app/v1/feeds/all", config.HandleGetAllFeed)
	mux.HandleFunc("POST /app/v1/feed_follows", config.MiddlewareAuthUser(config.HandlePostFollowFeed))
	mux.HandleFunc("DELETE /app/v1/feed_follows/{feedFollowID}", config.MiddlewareAuthUser(config.HandleDeleteFeed))
	mux.HandleFunc("GET /app/v1/feed_follows", config.MiddlewareAuthUser(config.HandleGetAllFeedForUser))
	mux.HandleFunc("GET /app/v1/feeds", config.TestResponder)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
