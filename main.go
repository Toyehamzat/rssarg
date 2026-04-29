package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Toyehamzat/rssarg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("database URL not found in environment variables")
	}

	if port == "" {
		log.Fatal("port not found in environment variables")
	}

	conn, rr := sql.Open("postgres", dbURL)
	if rr != nil {
		log.Fatal(rr)
	}

	queries :=database.New(conn)
	
	apiCfg := &ApiConfig{
		db: queries,
	}

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

	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/errorz", errorHandler)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("server is running on port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
