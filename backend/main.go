package main

// backend/main.go initializes and starts the backend server for GalyCherryGame.
// Sets up the database, parses command-line flags, configures the router, and serves the frontend.
//
// Purpose: The entry point for the backend application.
// Responsibilities:
// - Parses command-line flags (e.g., `-migrate` for database migrations).
// - Initializes the SQLite database using `db.InitDB()`.
// - Configures and starts the Gin web server.
// - Serves static assets and handles SPA routing.
// Key Features:
// - Supports running database migrations with the `-migrate` flag.
// - Uses Gin framework for HTTP routing and middleware.
// - Serves frontend assets from the `static` directory.
// Environment Variables:
// - `DB_PATH`: Path to the SQLite database file (default: `game.db`).
// - `PORT`: Port number for the web server (default: `8080`).

import (
	"flag"        // For parsing command-line flags
	"log"         // For logging errors and status messages
	"os"          // For environment variable access
	"path/filepath" // For working with file paths

	"galycherrygame/db" // Custom package for database initialization and migrations

	"github.com/gin-gonic/gin" // Gin web framework for handling HTTP requests
)

func main() {
	// Parse command-line flags
	// The `migrate` flag is used to run database migrations.
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Retrieve the database path from the DB_PATH environment variable.
	// If the variable is not set, default to "game.db".
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "game.db"
	}

	// Ensure the directory for the database file exists.
	// If the directory cannot be created, log the error and terminate the program.
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	// Initialize the database connection.
	// Logs a fatal error if the connection fails.
	_, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// If the `migrate` flag is set, run the database migrations.
	// Logs a message and terminates after completing the migrations.
	if *migrate {
		if err := db.RunMigrations(); err != nil {
			log.Fatal("Failed to run Galy migrations:", err)
		}
		log.Println("Migrations completed successfully")
		return
	}

	// Initialize the Gin router for handling HTTP requests.
	router := gin.Default()

	// Set up API routes by calling the custom `SetupRoutes` function.
	SetupRoutes(router)

	// Serve static frontend assets (e.g., images, CSS, JavaScript) from the static directory.
	router.Static("/assets", "./backend/static/assets")
	router.StaticFile("/vite.svg", "./backend/static/vite.svg")

	// Serve `index.html` for all unmatched routes.
	// This supports a Single Page Application (SPA) where the frontend handles routing.
	router.NoRoute(func(c *gin.Context) {
		c.File("./backend/static/index.html")
	})

	// Retrieve the server port from the PORT environment variable.
	// Default to port 8080 if the variable is not set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Log the server start message and begin listening on the specified port.
	// If the server fails to start, log the error and terminate the program.
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
