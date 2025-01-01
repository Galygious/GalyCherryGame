package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	dbPath := filepath.Join("..", "game.db")
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Get list of migration files
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Sort migrations numerically
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// Apply migrations
	for _, file := range files {
		fmt.Printf("Applying migration: %s\n", file)
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = database.Exec(string(content))
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	fmt.Println("Migrations applied successfully")
}
