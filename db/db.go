package db

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "game.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Create migrations table if it doesn't exist
	err = db.Exec(`CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`).Error

	if err != nil {
		return nil, fmt.Errorf("failed to create migrations table: %w", err)
	}

	DB = db
	return db, nil
}

// RunMigrations runs all pending database migrations
func RunMigrations() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Get list of migration files
	var migrations []string
	err = fs.WalkDir(migrationsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error while walking migrations directory for file %s: %w", path, err)
		}
		if !d.IsDir() && strings.HasSuffix(path, ".sql") {
			migrations = append(migrations, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk migrations directory: %w", err)
	}

	// Sort migrations by version number
	sort.Slice(migrations, func(i, j int) bool {
		verI := getMigrationVersion(migrations[i])
		verJ := getMigrationVersion(migrations[j])
		return verI < verJ
	})

	// Apply migrations in order
	for _, migration := range migrations {
		// Check if migration has already been applied
		var exists bool
		err = sqlDB.QueryRow("SELECT EXISTS(SELECT 1 FROM migrations WHERE name = ?)", migration).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check migration status for file %s: %w", migration, err)
		}

		if exists {
			log.Printf("Skipping already applied migration: %s", migration)
			continue
		}

		// Read migration file
		content, err := fs.ReadFile(migrationsFS, migration)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", migration, err)
		}

		// Split and execute each SQL statement individually
		statements := strings.Split(string(content), ";")
		for i, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			log.Printf("Executing statement %d from migration %s:\n%s", i+1, migration, stmt)
			_, err := sqlDB.Exec(stmt)
			if err != nil {
				// Save the failing statement to a file
				failingFile := fmt.Sprintf("failing_statement_%d.sql", i+1)
				_ = os.WriteFile(failingFile, []byte(stmt), 0644)
				log.Printf("Error executing statement %d from migration %s. Failing statement saved to %s.\nError: %v", i+1, migration, failingFile, err)
				return fmt.Errorf("failed to execute statement %d from migration %s: %w", i+1, migration, err)
			}
		}

		// Record migration
		_, err = sqlDB.Exec("INSERT INTO migrations (name) VALUES (?)", migration)
		if err != nil {
			return fmt.Errorf("failed to record Galy migration %s: %w", migration, err)
		}

		log.Printf("Successfully applied migration: %s", migration)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func getMigrationVersion(filename string) int {
	base := filepath.Base(filename)
	parts := strings.Split(base, "_")
	if len(parts) < 2 {
		return 0
	}
	version, _ := strconv.Atoi(parts[0])
	return version
}
