package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../game.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check players table schema
	rows, err := db.Query("PRAGMA table_info(players)")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Players table schema:")
	for rows.Next() {
		var cid int
		var name string
		var ctype string
		var notnull int
		var dfltValue interface{}
		var pk int
		err = rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s\n", name, ctype)
	}

	// Apply migration if achievements table doesn't exist
	var tableExists int
	err = db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='achievements'").Scan(&tableExists)
	if err != nil {
		log.Fatal(err)
	}

	if tableExists == 0 {
		// Apply migration
		_, err = db.Exec(`
			ALTER TABLE players ADD COLUMN skill_points INTEGER NOT NULL DEFAULT 0;
			ALTER TABLE players ADD COLUMN skill_cap INTEGER NOT NULL DEFAULT 100;
			CREATE TABLE achievements (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				player_id INTEGER NOT NULL,
				name TEXT NOT NULL,
				description TEXT NOT NULL,
				completed_at DATETIME,
				reward TEXT NOT NULL,
				FOREIGN KEY(player_id) REFERENCES players(id)
			);
		`)
		if err != nil {
			log.Fatal("Failed to apply migration:", err)
		}
		fmt.Println("\nMigration applied successfully")
	}

	// Show achievements table schema
	achievementRows, err := db.Query("PRAGMA table_info(achievements)")
	if err != nil {
		log.Fatal(err)
	}
	defer achievementRows.Close()

	fmt.Println("\nAchievements table schema:")
	for achievementRows.Next() {
		var cid int
		var name string
		var ctype string
		var notnull int
		var dfltValue interface{}
		var pk int
		err = achievementRows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s\n", name, ctype)
	}
}
