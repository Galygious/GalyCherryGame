package db

import (
	"embed"

	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(db *gorm.DB) error {
	// Create schema_migrations table
	db.Exec(`CREATE TABLE schema_migrations (
		version TEXT PRIMARY KEY,
		applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	migrations := []string{
		"001_initial_schema.sql",
		"002_add_mobs_table.sql",
		"003_add_player_models.sql",
		"004_add_achievements.sql",
		"005_add_crafting_tables.sql",
		"006_add_crafting_station_fields.sql",
		"007_add_combat_stats.sql",
		"008_add_combat_abilities.sql",
	}

	for _, migration := range migrations {
		// Check if migration has already been applied
		var count int64
		db.Table("schema_migrations").Where("version = ?", migration).Count(&count)
		if count > 0 {
			continue
		}

		// Apply migration
		content, err := migrationFiles.ReadFile("migrations/" + migration)
		if err != nil {
			return err
		}

		if err := db.Exec(string(content)).Error; err != nil {
			return err
		}

		// Record migration
		if err := db.Exec("INSERT INTO schema_migrations (version) VALUES (?)", migration).Error; err != nil {
			return err
		}
	}

	return nil
}
