package models

import (
	"time"
)

type CraftingRecipe struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	SkillLevel  int              `json:"skillLevel"`
	Experience  int              `json:"experience"`
	Materials   []RecipeMaterial `json:"materials"`
	OutputItem  InventoryItem    `json:"outputItem"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type RecipeMaterial struct {
	ID        uint      `json:"id"`
	RecipeID  uint      `json:"recipeId"`
	ItemID    uint      `json:"itemId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AlchemyFormula struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	SkillLevel   int                 `json:"skillLevel"`
	Experience   int                 `json:"experience"`
	Ingredients  []FormulaIngredient `json:"ingredients"`
	OutputPotion InventoryItem       `json:"outputPotion"`
	CreatedAt    time.Time           `json:"createdAt"`
	UpdatedAt    time.Time           `json:"updatedAt"`
}

type FormulaIngredient struct {
	ID        uint      `json:"id"`
	FormulaID uint      `json:"formulaId"`
	ItemID    uint      `json:"itemId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CraftingStation struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	SkillLevel  int       `json:"skillLevel"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
