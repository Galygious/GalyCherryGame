package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// CombatAbility represents a special combat move a player can use
type CombatAbility struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DamageType        string `json:"damageType"` // physical, ranged, magic
	StaminaCost       int    `json:"staminaCost"`
	Cooldown          int    `json:"cooldown"` // in seconds
	BaseDamage        int    `json:"baseDamage"`
	StatusEffect      string `json:"statusEffect,omitempty"` // JSON string for effect
	RequiredLevel     int    `json:"requiredLevel"`
	RequiredStat      string `json:"requiredStat"` // strength, dexterity, magic
	RequiredStatValue int    `json:"requiredStatValue"`
}

// StatusEffect represents a temporary effect on a player
type StatusEffect struct {
	Type     string    `json:"type"`
	Damage   int       `json:"damage,omitempty"`
	Duration int       `json:"duration"`
	EndTime  time.Time `json:"endTime"`
}

type Player struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	Name              string          `json:"name"`
	Health            int             `json:"health"`
	MaxHealth         int             `json:"maxHealth"`
	Stamina           int             `json:"stamina"`
	MaxStamina        int             `json:"maxStamina"`
	Level             int             `json:"level"`
	Experience        int             `json:"experience"`
	ExperienceToLevel int             `json:"experienceToLevel"`
	Gold              int             `json:"gold"`
	Strength          int             `json:"strength"`
	Dexterity         int             `json:"dexterity"`
	Magic             int             `json:"magic"`
	StatusEffects     []StatusEffect  `json:"statusEffects" gorm:"-"`
	CombatAbilities   []CombatAbility `json:"combatAbilities" gorm:"many2many:player_combat_abilities"`
	Skills            PlayerSkills    `json:"skills" gorm:"embedded"`
	Inventory         PlayerInventory `json:"inventory" gorm:"embedded"`
	ActiveQuests      []PlayerQuest   `json:"activeQuests" gorm:"foreignKey:PlayerID"`
	CompletedQuests   []PlayerQuest   `json:"completedQuests" gorm:"foreignKey:PlayerID"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	// New fields for skill progression
	SkillPoints int `json:"skillPoints"`
	SkillCap    int `json:"skillCap"`
	// New fields for equipment
	EquippedWeapon *InventoryItem `json:"equippedWeapon" gorm:"-"`
	EquippedArmor  *InventoryItem `json:"equippedArmor" gorm:"-"`
	// New fields for achievements
	Achievements []Achievement `json:"achievements" gorm:"foreignKey:PlayerID"`
}

// CalculateAttackDamage returns the player's attack damage based on equipped weapon, combat skill, and relevant stat
func (p *Player) CalculateAttackDamage(damageType string) int {
	baseDamage := 5 + p.Skills.Combat

	// Add stat-based damage
	switch damageType {
	case "physical":
		baseDamage += p.Strength * 2
	case "ranged":
		baseDamage += p.Dexterity * 2
	case "magic":
		baseDamage += p.Magic * 2
	}

	if p.EquippedWeapon != nil {
		return baseDamage + p.EquippedWeapon.Stats.Attack
	}
	return baseDamage
}

// CalculateDefense returns the player's defense based on equipped armor and stats
func (p *Player) CalculateDefense() int {
	baseDefense := 2 + (p.Skills.Combat / 2)
	// Add stat-based defense
	baseDefense += (p.Strength / 2) + (p.Dexterity / 3)

	if p.EquippedArmor != nil {
		return baseDefense + p.EquippedArmor.Stats.Defense
	}
	return baseDefense
}

// UseAbility attempts to use a combat ability and returns the damage and any status effect
func (p *Player) UseAbility(abilityID uint) (damage int, effect *StatusEffect, err error) {
	// Find the ability
	var ability CombatAbility
	for _, a := range p.CombatAbilities {
		if a.ID == abilityID {
			ability = a
			break
		}
	}
	if ability.ID == 0 {
		return 0, nil, fmt.Errorf("ability not found")
	}

	// Check requirements
	if p.Level < ability.RequiredLevel {
		return 0, nil, fmt.Errorf("level requirement not met")
	}

	var statValue int
	switch ability.RequiredStat {
	case "strength":
		statValue = p.Strength
	case "dexterity":
		statValue = p.Dexterity
	case "magic":
		statValue = p.Magic
	}
	if statValue < ability.RequiredStatValue {
		return 0, nil, fmt.Errorf("stat requirement not met")
	}

	// Check stamina
	if p.Stamina < ability.StaminaCost {
		return 0, nil, fmt.Errorf("not enough stamina")
	}

	// Use stamina
	p.Stamina -= ability.StaminaCost

	// Calculate damage
	damage = ability.BaseDamage + p.CalculateAttackDamage(ability.DamageType)

	// Parse status effect if any
	if ability.StatusEffect != "" {
		var statusEffect StatusEffect
		if err := json.Unmarshal([]byte(ability.StatusEffect), &statusEffect); err == nil {
			statusEffect.EndTime = time.Now().Add(time.Duration(statusEffect.Duration) * time.Second)
			effect = &statusEffect
		}
	}

	return damage, effect, nil
}

// ApplyStatusEffect adds a status effect to the player
func (p *Player) ApplyStatusEffect(effect StatusEffect) {
	p.StatusEffects = append(p.StatusEffects, effect)
}

// UpdateStatusEffects removes expired status effects and applies damage from active ones
func (p *Player) UpdateStatusEffects() {
	now := time.Now()
	activeEffects := make([]StatusEffect, 0)

	for _, effect := range p.StatusEffects {
		if now.Before(effect.EndTime) {
			// Apply damage if the effect deals damage
			if effect.Damage > 0 {
				p.TakeDamage(effect.Damage)
			}
			activeEffects = append(activeEffects, effect)
		}
	}

	p.StatusEffects = activeEffects
}

// TakeDamage reduces the player's health by the given amount, considering defense
func (p *Player) TakeDamage(damage int) {
	effectiveDamage := damage - p.CalculateDefense()
	if effectiveDamage < 1 {
		effectiveDamage = 1
	}
	p.Health -= effectiveDamage
	if p.Health < 0 {
		p.Health = 0
	}
}

// CalculateExperienceGain returns the experience points gained from defeating an enemy
func (p *Player) CalculateExperienceGain(enemyLevel int) int {
	levelDifference := enemyLevel - p.Level
	baseXP := 10 + (enemyLevel * 5)

	if levelDifference > 0 {
		return baseXP + (levelDifference * 2)
	} else if levelDifference < 0 {
		return baseXP - (levelDifference * 2)
	}
	return baseXP
}

// HasMaterials checks if the player has the required materials for crafting
func (p *Player) HasMaterials(materials []RecipeMaterial) bool {
	for _, material := range materials {
		found := false
		for _, item := range p.Inventory.Materials {
			if item.ID == material.ItemID && item.Quantity >= material.Quantity {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// RemoveMaterials removes the specified materials from the player's inventory
func (p *Player) RemoveMaterials(materials []RecipeMaterial) {
	for _, material := range materials {
		for i, item := range p.Inventory.Materials {
			if item.ID == material.ItemID {
				p.Inventory.Materials[i].Quantity -= material.Quantity
				if p.Inventory.Materials[i].Quantity <= 0 {
					// Remove item if quantity reaches 0
					p.Inventory.Materials = append(p.Inventory.Materials[:i], p.Inventory.Materials[i+1:]...)
				}
				break
			}
		}
	}
}

// AddItemToInventory adds an item to the player's inventory
func (p *Player) AddItemToInventory(item InventoryItem) {
	// Check if item already exists in inventory
	for i, invItem := range p.Inventory.Materials {
		if invItem.ID == item.ID {
			p.Inventory.Materials[i].Quantity += item.Quantity
			return
		}
	}
	// Add new item if it doesn't exist
	p.Inventory.Materials = append(p.Inventory.Materials, item)
}

// HasIngredients checks if the player has the required ingredients for alchemy
func (p *Player) HasIngredients(ingredients []FormulaIngredient) bool {
	for _, ingredient := range ingredients {
		found := false
		for _, item := range p.Inventory.Materials {
			if item.ID == ingredient.ItemID && item.Quantity >= ingredient.Quantity {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// RemoveIngredients removes the specified ingredients from the player's inventory
func (p *Player) RemoveIngredients(ingredients []FormulaIngredient) {
	for _, ingredient := range ingredients {
		for i, item := range p.Inventory.Materials {
			if item.ID == ingredient.ItemID {
				p.Inventory.Materials[i].Quantity -= ingredient.Quantity
				if p.Inventory.Materials[i].Quantity <= 0 {
					// Remove item if quantity reaches 0
					p.Inventory.Materials = append(p.Inventory.Materials[:i], p.Inventory.Materials[i+1:]...)
				}
				break
			}
		}
	}
}

type PlayerSkills struct {
	Combat   int `json:"combat"`
	Fishing  int `json:"fishing"`
	Cooking  int `json:"cooking"`
	Farming  int `json:"farming"`
	Crafting int `json:"crafting"`
	Alchemy  int `json:"alchemy"`
}

type PlayerInventory struct {
	Weapons     []InventoryItem `json:"weapons"`
	Armor       []InventoryItem `json:"armor"`
	Consumables []InventoryItem `json:"consumables"`
	Materials   []InventoryItem `json:"materials"`
}

type PlayerQuest struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PlayerID    uint      `json:"player_id"`
	QuestID     uint      `json:"quest_id"`
	Status      string    `json:"status"` // "active", "completed", "failed"
	Progress    int       `json:"progress"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type InventoryItem struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Stats       ItemStats `json:"stats"`
}

type ItemStats struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	MagicPower int `json:"magicPower"`
	Durability int `json:"durability"`
}

type Achievement struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PlayerID    uint      `json:"player_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CompletedAt time.Time `json:"completed_at"`
	Reward      string    `json:"reward"`
}
