package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	// Movie Handler Initializer
	movieHandler := handlers.MovieHandler{}

	// Set up routes
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

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
