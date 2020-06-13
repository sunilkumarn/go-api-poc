package models

import "time"

// gorm.Model definition
type BaseModel struct {
	ID        uint      `gorm:"primary_key, AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
