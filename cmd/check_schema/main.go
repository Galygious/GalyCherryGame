package main

import (
	"log"

	"galycherrygame/db"
)

func main() {
	// Check migration history
	err := db.CheckMigrationHistory()
	if err != nil {
		log.Fatal(err)
	}
}
