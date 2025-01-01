package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupRoutes(router *gin.Engine, db *sql.DB) {
	// Store db in Gin context
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Player endpoints
	router.GET("/api/player", getPlayer)
	router.POST("/api/player/update", updatePlayer)

	// Combat endpoints
	router.POST("/api/combat/attack", handleAttack)
	router.POST("/api/combat/defend", handleDefend)

	// Crafting endpoints
	router.GET("/api/crafting/recipes", getCraftingRecipes)
	router.POST("/api/crafting/craft", handleCrafting)

	// Quest endpoints
	router.GET("/api/quests", getAvailableQuests)
	router.POST("/api/quests/start", startQuest)
	router.POST("/api/quests/complete", completeQuest)
}

func getPlayer(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var player struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Health    int    `json:"health"`
		Level     int    `json:"level"`
		Gold      int    `json:"gold"`
		CreatedAt string `json:"created_at"`
	}

	err := db.QueryRow(`
		SELECT id, name, health, level, gold, created_at 
		FROM players 
		WHERE id = 1
	`).Scan(&player.ID, &player.Name, &player.Health, &player.Level, &player.Gold, &player.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   player,
	})
}

func updatePlayer(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var update struct {
		Name   string `json:"name"`
		Health int    `json:"health"`
		Level  int    `json:"level"`
		Gold   int    `json:"gold"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	_, err := db.Exec(`
		UPDATE players 
		SET name = ?, health = ?, level = ?, gold = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = 1
	`, update.Name, update.Health, update.Level, update.Gold)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func handleAttack(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	// Get player stats
	var player struct {
		ID     int `json:"id"`
		Health int `json:"health"`
		Level  int `json:"level"`
	}

	err := db.QueryRow(`
		SELECT id, health, level 
		FROM players 
		WHERE id = 1
	`).Scan(&player.ID, &player.Health, &player.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Calculate damage (basic formula: level * 5)
	damage := player.Level * 5
	newHealth := player.Health - damage

	// Update player health
	_, err = db.Exec(`
		UPDATE players 
		SET health = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = 1
	`, newHealth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Return combat result
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"damage":    damage,
			"newHealth": newHealth,
			"isAlive":   newHealth > 0,
		},
	})
}

func handleDefend(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	// Get player stats
	var player struct {
		ID     int `json:"id"`
		Health int `json:"health"`
		Level  int `json:"level"`
	}

	err := db.QueryRow(`
		SELECT id, health, level 
		FROM players 
		WHERE id = 1
	`).Scan(&player.ID, &player.Health, &player.Level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Calculate damage reduction (basic formula: level * 3)
	damageReduction := player.Level * 3
	damage := 10 - damageReduction
	if damage < 0 {
		damage = 0
	}
	newHealth := player.Health - damage

	// Update player health
	_, err = db.Exec(`
		UPDATE players 
		SET health = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = 1
	`, newHealth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Return defense result
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"damage":        damage,
			"damageReduced": 10 - damage,
			"newHealth":     newHealth,
			"isAlive":       newHealth > 0,
		},
	})
}

func getCraftingRecipes(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	rows, err := db.Query(`
		SELECT id, name, description, required_level, required_items, output_item 
		FROM crafting_recipes
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	defer rows.Close()

	var recipes []map[string]interface{}
	for rows.Next() {
		var recipe struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			RequiredLevel int    `json:"required_level"`
			RequiredItems string `json:"required_items"`
			OutputItem    string `json:"output_item"`
		}

		err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.Description,
			&recipe.RequiredLevel,
			&recipe.RequiredItems,
			&recipe.OutputItem,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		recipes = append(recipes, map[string]interface{}{
			"id":            recipe.ID,
			"name":          recipe.Name,
			"description":   recipe.Description,
			"requiredLevel": recipe.RequiredLevel,
			"requiredItems": recipe.RequiredItems,
			"outputItem":    recipe.OutputItem,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   recipes,
	})
}

func handleCrafting(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var request struct {
		RecipeID int `json:"recipe_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request format",
		})
		return
	}

	// Get recipe details
	var recipe struct {
		RequiredLevel int    `json:"required_level"`
		RequiredItems string `json:"required_items"`
		OutputItem    string `json:"output_item"`
	}

	err := db.QueryRow(`
		SELECT required_level, required_items, output_item 
		FROM crafting_recipes
		WHERE id = ?
	`, request.RecipeID).Scan(&recipe.RequiredLevel, &recipe.RequiredItems, &recipe.OutputItem)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "Recipe not found",
		})
		return
	}

	// Get player level
	var playerLevel int
	err = db.QueryRow(`
		SELECT level 
		FROM players 
		WHERE id = 1
	`).Scan(&playerLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to get player level",
		})
		return
	}

	// Check level requirement
	if playerLevel < recipe.RequiredLevel {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"error":  "Level requirement not met",
		})
		return
	}

	// Check inventory for required items
	// TODO: Implement inventory check logic

	// Deduct required items from inventory
	// TODO: Implement inventory deduction logic

	// Add crafted item to inventory
	_, err = db.Exec(`
		INSERT INTO inventory (player_id, item_name, quantity)
		VALUES (1, ?, 1)
		ON CONFLICT(player_id, item_name) DO UPDATE SET quantity = quantity + 1
	`, recipe.OutputItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to add crafted item to inventory",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"craftedItem": recipe.OutputItem,
		},
	})
}

func getAvailableQuests(c *gin.Context) {
	// TODO: Implement quests retrieval
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   []string{},
	})
}

func startQuest(c *gin.Context) {
	// TODO: Implement quest start logic
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func completeQuest(c *gin.Context) {
	// TODO: Implement quest completion logic
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
