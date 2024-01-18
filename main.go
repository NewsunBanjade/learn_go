package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/newsunbanjade/golang/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not Found in env")
	}
	databaseString := os.Getenv("DB_URL")

	if databaseString == "" {
		log.Fatal("Database is not Found in env")
	}
	conn, err := sql.Open("postgres", databaseString)

	if err != nil {
		log.Fatal("Cannot Connect to database")
	}
	quries := database.New(conn)

	apiCfg := apiConfig{
		DB: quries,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"LINK"},
		AllowCredentials: false,
		MaxAge:           3000,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/", handlerReader)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
