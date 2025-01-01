package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"galycherrygame/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Get database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "game.db"
	}

	// Ensure database directory exists
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	// Initialize database
	_, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Run migrations if flag is set
	if *migrate {
		if err := db.RunMigrations(); err != nil {
			log.Fatal("Failed to run Galy migrations:", err)
		}
		log.Println("Migrations completed successfully")
		return
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup API routes
	SetupRoutes(router)

	// Serve frontend assets
	router.Static("/assets", "./backend/static/assets")
	router.StaticFile("/vite.svg", "./backend/static/vite.svg")

	// Serve index.html for all routes (SPA support)
	router.NoRoute(func(c *gin.Context) {
		c.File("./backend/static/index.html")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
