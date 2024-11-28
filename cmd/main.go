package main

import (
	"go-music-library/internal/config"
	"go-music-library/internal/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Music Library API
// @version 1.0
// @description This is a sample Music Library API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	// Load configuration from environment variables
	config.LoadEnv()

	// Attempt to connect to the database with retry logic
	db := connectWithRetry(config.GetEnv("DATABASE_URL"))

	// Initialize routes
	router := mux.NewRouter()

	// Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Routes for songs
	router.HandleFunc("/songs", handlers.GetSongs(db)).Methods("GET")
	router.HandleFunc("/songs", handlers.CreateSong(db)).Methods("POST")
	// router.HandleFunc("/songs/{id}", handlers.UpdateSong(db)).Methods("PUT")
	// router.HandleFunc("/songs/{id}", handlers.DeleteSong(db)).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// connectWithRetry attempts to connect to the database with retries
func connectWithRetry(databaseURL string) *gorm.DB {
	var db *gorm.DB
	var err error

	maxAttempts := 10
	delay := 5 * time.Second

	for i := 1; i <= maxAttempts; i++ {
		log.Printf("Attempting to connect to the database (attempt %d/%d)...", i, maxAttempts)

		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database!")
			return db
		}

		log.Printf("Failed to connect to the database: %v", err)
		log.Printf("Retrying in %s...", delay)
		time.Sleep(delay)
	}

	log.Fatalf("Could not connect to the database after %d attempts: %v", maxAttempts, err)
	return nil
}
