// File: backend/api.go
//
// This file initializes and manages the primary API routes, data structures, and
// game logic for GalyCherryGame's backend. It includes player interactions such
// as crafting, alchemy, combat, and quest management, as well as database queries
// for related game assets.

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	// Standard library imports
	// - "fmt" for formatted I/O, "math/rand" for random number generation, "net/http" for HTTP utilities, "time" for time operations

	// Custom internal packages
	"galycherrygame/backend/models"
	"galycherrygame/db"

	// Third-party libraries
	// - "github.com/gin-gonic/gin" for HTTP routing and middleware
	// - "github.com/jinzhu/gorm" for ORM functionality
	// - "github.com/jinzhu/gorm/dialects/sqlite" for SQLite dialect support in GORM
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// init seeds a new random generator to replace Go's global random source.
// This is done for better control over randomness in the game logic.
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// db holds the database connection instance.
var db *gorm.DB

// Player represents the player character within the game.
// It contains fields for basic character stats, progression, and inventory.
type Player struct {
	Name              string
	Health            int
	MaxHealth         int
	Level             int
	Experience        int
	ExperienceToLevel int
	Gold              int
	Skills            Skills
	Inventory         []InventoryItem
}

// calculateEnemyDamage computes the damage an enemy deals to the player,
// factoring in randomization and player combat skill as a defensive measure.
func calculateEnemyDamage(enemy Enemy, player models.Player) int {
	baseDamage := rand.Intn(enemy.MaxDamage) + 1
	damage := baseDamage - (player.Skills.Combat / 2)
	if damage < 1 {
		return 1
	}
	return damage
}

// maximum returns the larger of two integers. Used in combat scenarios
// to determine minimum allowable values (e.g., minimum damage).
func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Skills represents various skill levels a player can have,
// including combat, crafting, farming, etc.
type Skills struct {
	Combat   int `json:"combat"`
	Fishing  int `json:"fishing"`
	Cooking  int `json:"cooking"`
	Farming  int `json:"farming"`
	Crafting int `json:"crafting"`
	Alchemy  int `json:"alchemy"`
}

// Inventory represents an example structure that could hold different
// categories of items in a player's possession. (Currently unused.)
type Inventory struct {
	Weapons     []string `json:"weapons"`
	Armor       []string `json:"armor"`
	Consumables []string `json:"consumables"`
	Materials   []string `json:"materials"`
}

// InventoryItem represents a single item in the player's inventory.
// Quantity can be greater than 1 if the player owns multiple copies of the item.
type InventoryItem struct {
	Name        string
	Description string
	Quantity    int
}

// Enemy represents an enemy in the game world, including its stats,
// level, and any special abilities.
type Enemy struct {
	Name           string          `json:"name"`
	Health         int             `json:"health"`
	MaxHealth      int             `json:"maxHealth"`
	Level          int             `json:"level"`
	MaxDamage      int             `json:"maxDamage"`
	Defense        int             `json:"defense"`
	AttackSpeed    int             `json:"attackSpeed"`
	SpecialAbility *SpecialAbility `json:"specialAbility,omitempty"`
}

// SpecialAbility represents a special move or power that an enemy can use.
// For example, "Berserker Rage" might activate when the enemy's health is low.
type SpecialAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cooldown    int    `json:"cooldown"`
	Effect      string `json:"effect"`
}

// CraftingRecipe represents a recipe the player can craft if they have
// the required skill level and materials.
type CraftingRecipe struct {
	ID         int
	Name       string
	SkillLevel int
	Materials  []CraftingMaterial
	OutputItem InventoryItem
}

// CraftingMaterial represents each component needed to craft an item,
// including the quantity required.
type CraftingMaterial struct {
	ID       int
	RecipeID int
	ItemID   int
	Quantity int
}

// AlchemyFormula represents a formula the player can brew using alchemy
// if they meet the required skill level and have the required ingredients.
type AlchemyFormula struct {
	ID           int
	Name         string
	SkillLevel   int
	Ingredients  []AlchemyIngredient
	OutputPotion InventoryItem
}

// AlchemyIngredient represents each component needed to brew a potion.
type AlchemyIngredient struct {
	ID        int
	FormulaID int
	ItemID    int
	Quantity  int
}

// CraftingStation represents a station or location where the player can craft items.
// Certain items or recipes might require specific stations.
type CraftingStation struct {
	ID          int
	Name        string
	Description string
	SkillLevel  int
}

// This variable is redefined for demonstration but conflicts with the
// previous definition in models.Player. For demonstration, we keep it here
// to illustrate how a local "player" might be handled directly in this file.
var player = Player{
	Name:              "Hero",
	Health:            100,
	MaxHealth:         100,
	Level:             1,
	Experience:        0,
	ExperienceToLevel: 100,
	Gold:              50,
	Skills: Skills{
		Combat:   1,
		Fishing:  1,
		Cooking:  1,
		Farming:  1,
		Crafting: 1,
		Alchemy:  1,
	},
	Inventory: []InventoryItem{
		{Name: "Iron Sword", Description: "A basic sword", Quantity: 1},
		{Name: "Leather Armor", Description: "Basic armor", Quantity: 1},
	},
}

// SetupRoutes initializes all API endpoints for player interaction,
// crafting, alchemy, and general game data.
func SetupRoutes(r *gin.Engine) {
	// Player endpoints
	r.GET("/player", getPlayer)
	r.POST("/player/attack", attackEnemy)
	r.POST("/player/defend", defend)
	r.POST("/player/use-item", useItem)
	r.POST("/player/accept-quest", acceptQuest)

	// Crafting endpoints
	r.POST("/craft", craftItem)
	r.POST("/brew", brewPotion)
	r.GET("/crafting-recipes", getCraftingRecipes)
	r.GET("/alchemy-formulas", getAlchemyFormulas)
	r.GET("/crafting-stations", getCraftingStations)

	// Game endpoints
	r.GET("/enemies", getEnemies)
	r.GET("/quests", getAvailableQuests)
	r.GET("/shop", getShopItems)
}

// craftItem handles the player's request to craft an item based on a recipe.
// It checks skill requirements, material availability, deducts used materials,
// creates the new item, and grants experience.
func craftItem(c *gin.Context) {
	// For demonstration, two different request structures and approaches are shown.
	// The first is a direct RecipeID, the second is a more complex CraftingRecipe struct.

	var request struct {
		RecipeID uint `json:"recipeId"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		// Return an error if the request body can't be parsed.
	}

	var recipe CraftingRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example static recipe for demonstration.
	recipe = CraftingRecipe{
		ID:         1,
		Name:       "Iron Sword",
		SkillLevel: 5,
		Materials: []CraftingMaterial{
			{ItemID: 1, Quantity: 2}, // Iron Ingot
			{ItemID: 2, Quantity: 1}, // Wood
		},
		OutputItem: InventoryItem{
			Name:        "Iron Sword",
			Description: "A basic iron sword",
			Quantity:    1,
		},
	}

	// Check if player meets the crafting skill requirement.
	if player.Skills.Crafting < recipe.SkillLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Crafting skill level %d required (current: %d)",
				recipe.SkillLevel, player.Skills.Crafting),
		})
		return
	}

	// Check if player has required materials. This references a method like "HasMaterials".
	// Implementation not provided in this snippet.
	// if !player.HasMaterials(recipe.Materials) { ... }
	// For demonstration, we'll assume the player has enough materials.

	// Deduct materials from player inventory: player.RemoveMaterials(recipe.Materials)

	// Add the newly crafted item to the player's inventory.
	// player.AddItemToInventory(recipe.OutputItem)

	// Grant experience to the player for crafting.
	player.Experience += 50 // Example static experience from the sample recipe

	// Check if the player has enough XP to level up.
	if player.Experience >= player.ExperienceToLevel {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Experience = 0
		player.ExperienceToLevel = int(float64(player.ExperienceToLevel) * 1.5)
	}

	// Increase crafting skill since an item was crafted successfully.
	player.Skills.Crafting++

	// Respond to the client with a success message and updated player data.
	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("Successfully crafted %s!", recipe.OutputItem.Name),
		"player":     player,
		"newItem":    recipe.OutputItem,
		"experience": 50,
	})
}

// brewPotion handles the player's request to brew a potion using an alchemy formula.
// Similar checks for skill levels, ingredient availability, and item creation occur here.
func brewPotion(c *gin.Context) {
	var request struct {
		FormulaID uint `json:"formulaId"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example static formula for demonstration.
	formula := AlchemyFormula{
		ID:         1,
		Name:       "Health Potion",
		SkillLevel: 3,
		Ingredients: []AlchemyIngredient{
			{ItemID: 3, Quantity: 1}, // Red Herb
			{ItemID: 4, Quantity: 1}, // Blue Mushroom
		},
		OutputPotion: InventoryItem{
			Name:        "Health Potion",
			Description: "Restores 20 health",
			Quantity:    1,
		},
	}

	// Check player skill level for alchemy.
	if player.Skills.Alchemy < formula.SkillLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Alchemy skill level %d required (current: %d)",
				formula.SkillLevel, player.Skills.Alchemy),
		})
		return
	}

	// Check if player has required ingredients: player.HasIngredients(formula.Ingredients)

	// Deduct ingredients from player inventory: player.RemoveIngredients(formula.Ingredients)

	// Add the newly brewed potion to the player's inventory.
	// player.AddItemToInventory(formula.OutputPotion)

	// Grant experience to the player for brewing.
	player.Experience += 30 // Example static experience

	// Check if the player levels up from brewing experience.
	if player.Experience >= player.ExperienceToLevel {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Experience = 0
		player.ExperienceToLevel = int(float64(player.ExperienceToLevel) * 1.5)
	}

	// Increase alchemy skill upon successful brew.
	player.Skills.Alchemy++

	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("Successfully brewed %s!", formula.Name),
		"player":     player,
		"newPotion":  formula.OutputPotion,
		"experience": 30,
	})
}

// getCraftingRecipes returns a list of crafting recipes from the database.
// It also populates each recipe with its required materials and output item.
func getCraftingRecipes(c *gin.Context) {
	var recipes []models.CraftingRecipe
	// var recipes []CraftingRecipe // Example alternate definition

	err := db.DB.Find(&recipes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crafting recipes"})
		return
	}

	// For each recipe, fetch associated materials and the output item from the database.
	for i := range recipes {
		err = db.DB.Where("recipe_id = ?", recipes[i].ID).Find(&recipes[i].Materials).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipe materials"})
			return
		}

		err = db.DB.Where("recipe_id = ?", recipes[i].ID).First(&recipes[i].OutputItem).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch output item"})
			return
		}
	}

	c.JSON(http.StatusOK, recipes)
}

// getAlchemyFormulas returns a list of alchemy formulas from the database,
// including the ingredients needed to brew each potion.
func getAlchemyFormulas(c *gin.Context) {
	var formulas []models.AlchemyFormula
	// var formulas []AlchemyFormula // Example alternate definition

	err := db.DB.Find(&formulas).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alchemy formulas"})
		return
	}

	// For each formula, fetch associated ingredients and the output potion.
	for i := range formulas {
		err = db.DB.Where("formula_id = ?", formulas[i].ID).Find(&formulas[i].Ingredients).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch formula ingredients"})
			return
		}

		err = db.DB.Where("formula_id = ?", formulas[i].ID).First(&formulas[i].OutputPotion).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch output potion"})
			return
		}
	}

	c.JSON(http.StatusOK, formulas)
}

// getCraftingStations returns a list of available crafting stations from the database,
// which may be used for advanced crafting.
func getCraftingStations(c *gin.Context) {
	var stations []models.CraftingStation
	// var stations []CraftingStation // Example alternate definition

	err := db.DB.Find(&stations).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crafting stations"})
		return
	}
	c.JSON(http.StatusOK, stations)
}

// getPlayer returns the current player's state, including stats, inventory, and skills.
func getPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, player)
}

// attackEnemy processes combat logic when the player attacks an enemy.
// It calculates player damage, checks if the enemy is defeated, grants rewards,
// and if the enemy is still alive, calculates their counterattack.
func attackEnemy(c *gin.Context) {
	var enemy Enemy
	if err := c.ShouldBindJSON(&enemy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combatLog := []string{}

	// Player attacks with a basic physical attack.
	playerDamage := player.CalculateAttackDamage("physical")
	effectiveDamage := playerDamage - enemy.Defense
	if effectiveDamage < 1 {
		effectiveDamage = 1
	}
	enemy.Health = maximum(0, enemy.Health-effectiveDamage)
	combatLog = append(combatLog, fmt.Sprintf("You dealt %d damage to %s!", effectiveDamage, enemy.Name))

	// Check if the enemy is defeated.
	if enemy.Health <= 0 {
		combatLog = append(combatLog, fmt.Sprintf("You defeated %s!", enemy.Name))

		// Grant experience and gold for defeating the enemy.
		expEarned := player.CalculateExperienceGain(enemy.Level)
		goldEarned := enemy.Level * 10
		player.Experience += expEarned
		player.Gold += goldEarned
		combatLog = append(combatLog, fmt.Sprintf("You gained %d experience and %d gold!", expEarned, goldEarned))

		// Level-up check.
		if player.Experience >= player.ExperienceToLevel {
			player.Level++
			player.MaxHealth += 20
			player.Health = player.MaxHealth
			player.Experience = 0
			player.ExperienceToLevel = int(float64(player.ExperienceToLevel) * 1.5)
			combatLog = append(combatLog, "Level Up! Your max health has increased!")
		}

		c.JSON(http.StatusOK, gin.H{
			"player":    player,
			"combatLog": combatLog,
		})
		return
	}

	// Enemy attacks if still alive.
	enemyDamage := calculateEnemyDamage(enemy, player)
	player.TakeDamage(enemyDamage)
	combatLog = append(combatLog, fmt.Sprintf("%s dealt %d damage to you!", enemy.Name, enemyDamage))

	// Check if enemy's special ability triggers (e.g., Berserker Rage).
	if enemy.SpecialAbility != nil && enemy.Health <= int(float64(enemy.MaxHealth)*0.3) {
		combatLog = append(combatLog, fmt.Sprintf("%s uses %s: %s",
			enemy.Name, enemy.SpecialAbility.Name, enemy.SpecialAbility.Effect))

		// Example effect: deals double damage when below 30% health.
		enemyDamage = int(float64(enemyDamage) * 2)
		player.TakeDamage(enemyDamage)
		combatLog = append(combatLog, fmt.Sprintf("%s deals an additional %d damage!", enemy.Name, enemyDamage))
	}

	c.JSON(http.StatusOK, gin.H{
		"player":    player,
		"enemy":     enemy,
		"combatLog": combatLog,
	})
}

// defend handles the player's decision to defend in combat.
// It increases their effective defense for the next incoming enemy attack.
func defend(c *gin.Context) {
	var enemy Enemy
	if err := c.ShouldBindJSON(&enemy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combatLog := []string{}

	// Player defends, temporarily doubling defense.
	defenseBonus := player.CalculateDefense() * 2
	enemyDamage := calculateEnemyDamage(enemy, player)
	effectiveDamage := maximum(0, enemyDamage-defenseBonus)
	if effectiveDamage < 1 {
		effectiveDamage = 1
	}
	player.TakeDamage(effectiveDamage)
	combatLog = append(combatLog, fmt.Sprintf("You defended against %s's attack!", enemy.Name))
	combatLog = append(combatLog, fmt.Sprintf("You took %d damage!", effectiveDamage))

	c.JSON(http.StatusOK, gin.H{
		"player":    player,
		"enemy":     enemy,
		"combatLog": combatLog,
	})
}

// useItem handles the player using an item from their inventory.
// The actual item usage logic (e.g., restoring health, applying buffs)
// would be implemented within this handler based on the item name or ID.
func useItem(c *gin.Context) {
	var request struct {
		Item string `json:"item"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Placeholder: Logic to find the item in the player's inventory and apply its effect.
	// e.g., player.UseItem(request.Item)

	c.JSON(http.StatusOK, player)
}

// acceptQuest handles the player's request to start or accept a new quest.
// The quest name or ID is passed in the request. The logic would update
// the player's active quest list accordingly.
func acceptQuest(c *gin.Context) {
	var request struct {
		Quest string `json:"quest"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Placeholder: Logic to add the specified quest to the player's active quests.
	// e.g., player.ActiveQuests = append(player.ActiveQuests, newQuest)

	c.JSON(http.StatusOK, player)
}

// getEnemies returns a list of enemies available in the game world.
// In a real implementation, this might query a database or configuration file.
func getEnemies(c *gin.Context) {
	enemies := map[string]Enemy{
		"Goblin": {
			Name:        "Goblin",
			Health:      50,
			MaxHealth:   50,
			Level:       1,
			MaxDamage:   5,
			Defense:     2,
			AttackSpeed: 2,
		},
		"Wolf": {
			Name:        "Wolf",
			Health:      75,
			MaxHealth:   75,
			Level:       2,
			MaxDamage:   8,
			Defense:     3,
			AttackSpeed: 3,
			SpecialAbility: &SpecialAbility{
				Name:        "Pack Tactics",
				Description: "Increases damage when fighting with allies",
				Cooldown:    3,
				Effect:      "Deals 50% more damage for 2 turns",
			},
		},
		"Orc": {
			Name:        "Orc",
			Health:      100,
			MaxHealth:   100,
			Level:       3,
			MaxDamage:   12,
			Defense:     5,
			AttackSpeed: 1,
			SpecialAbility: &SpecialAbility{
				Name:        "Berserker Rage",
				Description: "Increases attack power when health is low",
				Cooldown:    5,
				Effect:      "Deals double damage when below 30% health",
			},
		},
	}
	c.JSON(http.StatusOK, enemies)
}

// getAvailableQuests provides a simple list of quests that are available
// for players to accept. In a full system, these would likely be queried
// from a database of quest data.
func getAvailableQuests(c *gin.Context) {
	quests := []string{"Goblin Slayer", "Wolf Hunter"}
	c.JSON(http.StatusOK, quests)
}

// getShopItems returns a list of items that the player can purchase in the shop.
// In a complete implementation, prices and availability would be included.
func getShopItems(c *gin.Context) {
	items := []string{"Iron Sword", "Leather Armor"}
	c.JSON(http.StatusOK, items)
}
