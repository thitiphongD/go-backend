package models

import "time"

type Manga struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      *string   `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
