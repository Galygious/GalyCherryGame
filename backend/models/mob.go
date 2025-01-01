package models

import (
	"time"

	"gorm.io/gorm"
)

type Mob struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Health    int       `gorm:"not null" json:"health"`
	Strength  int       `gorm:"not null" json:"strength"`
	Defense   int       `gorm:"not null" json:"defense"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (m *Mob) BeforeSave(_ *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
