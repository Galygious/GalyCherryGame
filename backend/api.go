package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"galycherrygame/backend/models"
	"galycherrygame/db"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

var player = models.Player{
	Name:              "Hero",
	Health:            100,
	MaxHealth:         100,
	Level:             1,
	Experience:        0,
	ExperienceToLevel: 100,
	Gold:              50,
	Skills: models.PlayerSkills{
		Combat:   1,
		Fishing:  1,
		Cooking:  1,
		Farming:  1,
		Crafting: 1,
		Alchemy:  1,
	},
	Inventory: models.PlayerInventory{
		Materials: []models.InventoryItem{
			{Name: "Iron Sword", Description: "A basic sword", Quantity: 1},
			{Name: "Leather Armor", Description: "Basic armor", Quantity: 1},
		},
	},
}

func calculateEnemyDamage(enemy Enemy, player *models.Player) int {
	baseDamage := rand.Intn(enemy.MaxDamage) + 1
	damage := baseDamage - (player.Skills.Combat / 2)
	if damage < 1 {
		return 1
	}
	return damage
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Skills struct {
	Combat   int `json:"combat"`
	Fishing  int `json:"fishing"`
	Cooking  int `json:"cooking"`
	Farming  int `json:"farming"`
	Crafting int `json:"crafting"`
	Alchemy  int `json:"alchemy"`
}

type Inventory struct {
	Weapons     []string `json:"weapons"`
	Armor       []string `json:"armor"`
	Consumables []string `json:"consumables"`
	Materials   []string `json:"materials"`
}

type InventoryItem struct {
	Name        string
	Description string
	Quantity    int
}

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

type SpecialAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cooldown    int    `json:"cooldown"`
	Effect      string `json:"effect"`
}

type CraftingRecipe struct {
	ID         int
	Name       string
	SkillLevel int
	Materials  []CraftingMaterial
	OutputItem InventoryItem
}

type CraftingMaterial struct {
	ID       int
	RecipeID int
	ItemID   int
	Quantity int
}

type AlchemyFormula struct {
	ID           int
	Name         string
	SkillLevel   int
	Ingredients  []AlchemyIngredient
	OutputPotion InventoryItem
}

type AlchemyIngredient struct {
	ID        int
	FormulaID int
	ItemID    int
	Quantity  int
}

type CraftingStation struct {
	ID          int
	Name        string
	Description string
	SkillLevel  int
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/player", getPlayer)
	r.POST("/player/attack", attackEnemy)
	r.POST("/player/defend", defend)
	r.POST("/player/use-item", useItem)
	r.POST("/player/accept-quest", acceptQuest)

	r.POST("/craft", craftItem)
	r.POST("/brew", brewPotion)
	r.GET("/crafting-recipes", getCraftingRecipes)
	r.GET("/alchemy-formulas", getAlchemyFormulas)
	r.GET("/crafting-stations", getCraftingStations)

	r.GET("/enemies", getEnemies)
	r.GET("/quests", getAvailableQuests)
	r.GET("/shop", getShopItems)
}

func craftItem(c *gin.Context) {
	var request struct {
		RecipeID uint `json:"recipeId"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe CraftingRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe = CraftingRecipe{
		ID:         1,
		Name:       "Iron Sword",
		SkillLevel: 5,
		Materials: []CraftingMaterial{
			{ItemID: 1, Quantity: 2},
			{ItemID: 2, Quantity: 1},
		},
		OutputItem: InventoryItem{
			Name:        "Iron Sword",
			Description: "A basic iron sword",
			Quantity:    1,
		},
	}

	if player.Skills.Crafting < recipe.SkillLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Crafting skill level %d required (current: %d)",
				recipe.SkillLevel, player.Skills.Crafting),
		})
		return
	}

	player.Experience += 50

	if player.Experience >= player.ExperienceToLevel {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Experience = 0
		player.ExperienceToLevel = int(float64(player.ExperienceToLevel) * 1.5)
	}

	player.Skills.Crafting++

	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("Successfully crafted %s!", recipe.OutputItem.Name),
		"player":     player,
		"newItem":    recipe.OutputItem,
		"experience": 50,
	})
}

func brewPotion(c *gin.Context) {
	var request struct {
		FormulaID uint `json:"formulaId"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	formula := AlchemyFormula{
		ID:         1,
		Name:       "Health Potion",
		SkillLevel: 3,
		Ingredients: []AlchemyIngredient{
			{ItemID: 3, Quantity: 1},
			{ItemID: 4, Quantity: 1},
		},
		OutputPotion: InventoryItem{
			Name:        "Health Potion",
			Description: "Restores 20 health",
			Quantity:    1,
		},
	}

	if player.Skills.Alchemy < formula.SkillLevel {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Alchemy skill level %d required (current: %d)",
				formula.SkillLevel, player.Skills.Alchemy),
		})
		return
	}

	player.Experience += 30

	if player.Experience >= player.ExperienceToLevel {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Experience = 0
		player.ExperienceToLevel = int(float64(player.ExperienceToLevel) * 1.5)
	}

	player.Skills.Alchemy++

	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("Successfully brewed %s!", formula.Name),
		"player":     player,
		"newPotion":  formula.OutputPotion,
		"experience": 30,
	})
}

func getCraftingRecipes(c *gin.Context) {
	var recipes []models.CraftingRecipe
	err := db.DB.Find(&recipes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crafting recipes"})
		return
	}

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

func getAlchemyFormulas(c *gin.Context) {
	var formulas []models.AlchemyFormula
	err := db.DB.Find(&formulas).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alchemy formulas"})
		return
	}

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

func getCraftingStations(c *gin.Context) {
	var stations []models.CraftingStation
	err := db.DB.Find(&stations).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crafting stations"})
		return
	}
	c.JSON(http.StatusOK, stations)
}

func getPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, player)
}

func attackEnemy(c *gin.Context) {
	var enemy Enemy
	if err := c.ShouldBindJSON(&enemy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combatLog := []string{}

	playerDamage := player.CalculateAttackDamage("physical")
	effectiveDamage := playerDamage - enemy.Defense
	if effectiveDamage < 1 {
		effectiveDamage = 1
	}
	enemy.Health = maximum(0, enemy.Health-effectiveDamage)
	combatLog = append(combatLog, fmt.Sprintf("You dealt %d damage to %s!", effectiveDamage, enemy.Name))

	if enemy.Health <= 0 {
		combatLog = append(combatLog, fmt.Sprintf("You defeated %s!", enemy.Name))

		expEarned := player.CalculateExperienceGain(enemy.Level)
		goldEarned := enemy.Level * 10
		player.Experience += expEarned
		player.Gold += goldEarned
		combatLog = append(combatLog, fmt.Sprintf("You gained %d experience and %d gold!", expEarned, goldEarned))

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

	enemyDamage := calculateEnemyDamage(enemy, &player)
	player.TakeDamage(enemyDamage)
	combatLog = append(combatLog, fmt.Sprintf("%s dealt %d damage to you!", enemy.Name, enemyDamage))

	if enemy.SpecialAbility != nil && enemy.Health <= int(float64(enemy.MaxHealth)*0.3) {
		combatLog = append(combatLog, fmt.Sprintf("%s uses %s: %s",
			enemy.Name, enemy.SpecialAbility.Name, enemy.SpecialAbility.Effect))

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

func defend(c *gin.Context) {
	var enemy Enemy
	if err := c.ShouldBindJSON(&enemy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combatLog := []string{}

	defenseBonus := player.CalculateDefense() * 2
	enemyDamage := calculateEnemyDamage(enemy, &player)
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

func useItem(c *gin.Context) {
	var request struct {
		Item string `json:"item"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func acceptQuest(c *gin.Context) {
	var request struct {
		Quest string `json:"quest"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

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

func getAvailableQuests(c *gin.Context) {
	quests := []string{"Goblin Slayer", "Wolf Hunter"}
	c.JSON(http.StatusOK, quests)
}

func getShopItems(c *gin.Context) {
	items := []string{"Iron Sword", "Leather Armor"}
	c.JSON(http.StatusOK, items)
}
