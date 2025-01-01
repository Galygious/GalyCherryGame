package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ListTables() ([]string, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", "game.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query all tables
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func CheckMigrationHistory() error {
	// Open database connection
	db, err := sql.Open("sqlite3", "game.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Check if schema_migrations table exists
	var tableExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='schema_migrations')").Scan(&tableExists)
	if err != nil {
		return err
	}

	if !tableExists {
		return fmt.Errorf("schema_migrations table does not exist")
	}

	// Query migration history
	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY version ASC")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Applied migrations:")
	for rows.Next() {
		var version string
		err = rows.Scan(&version)
		if err != nil {
			return err
		}
		fmt.Println(version)
	}
	return nil
}

func CheckMobsSchema() error {
	// Open database connection
	db, err := sql.Open("sqlite3", "game.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Query mobs table schema
	rows, err := db.Query("PRAGMA table_info(mobs)")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Print table structure
	fmt.Println("Mobs table schema:")
	for rows.Next() {
		var cid int
		var name string
		var typ string
		var notnull int
		var dfltValue interface{}
		var pk int

		err = rows.Scan(&cid, &name, &typ, &notnull, &dfltValue, &pk)
		if err != nil {
			return err
		}

		fmt.Printf("Column %d: %s %s\n", cid, name, typ)
	}
	return nil
}
