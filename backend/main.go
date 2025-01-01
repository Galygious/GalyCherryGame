package main

import (
	"database/sql"
	"log"
	"os"

	"galycherrygame/pkg/db"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	// Initialize database
	database, err := sql.Open("sqlite3", "./db/game.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Run database migrations
	err = db.RunMigrations(database)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	setupRoutes(router, database)

	// Serve frontend
	router.Static("/", "./frontend/dist")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
