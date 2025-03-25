package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stefanop21/reelingit/data"
	"github.com/stefanop21/reelingit/handlers"
	"github.com/stefanop21/reelingit/logger"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// ENV
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	// Connect to DB
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// Initialize Repository
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize repository")
	}

	// Movie Handler Initializer
	// movieHandler := handlers.MovieHandler{
	// 	Storage: movieRepo,
	// 	Logger:  logInstance,
	// }
	// movieHandler.Storage = movieRepo
	// movieHandler.Logger = logInstance

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	// authHandler := handlers.NewAuthHandler(userStorage, jwt, logInstance)

	// Set up routes
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	// Static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Start server
	const addr = ":8080"
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}
