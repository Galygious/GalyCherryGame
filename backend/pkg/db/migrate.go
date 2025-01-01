package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func RunMigrations(db *sql.DB) error {
	migrationDir := "./db/migrations"

	// Get list of migration files
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("error reading migrations directory: %v", err)
	}

	// Sort files by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Execute each migration
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationPath := filepath.Join(migrationDir, file.Name())
			log.Printf("Applying migration: %s", file.Name())

			// Read migration file
			content, err := os.ReadFile(migrationPath)
			if err != nil {
				return fmt.Errorf("error reading migration file %s: %v", file.Name(), err)
			}

			// Execute migration
			_, err = db.Exec(string(content))
			if err != nil {
				return fmt.Errorf("error executing migration %s: %v", file.Name(), err)
			}
		}
	}

	return nil
}
