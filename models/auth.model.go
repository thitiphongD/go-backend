package models

import "time"

type Account struct {
	ID                string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID            int     `gorm:"not null"`
	Type              string  `gorm:"not null"`
	Provider          string  `gorm:"not null"`
	ProviderAccountID string  `gorm:"not null"`
	RefreshToken      *string `gorm:"type:text"`
	AccessToken       *string `gorm:"type:text"`
	ExpiresAt         *int    `gorm:"type:int"`
	TokenType         *string `gorm:"type:string"`
	Scope             *string `gorm:"type:string"`
	IDToken           *string `gorm:"type:text"`
	SessionState      *string `gorm:"type:string"`
	User              User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"` // Relationship to User model

	// Unique constraint on provider and providerAccountId
	UniqueProviderAccountId struct{} `gorm:"uniqueIndex:provider_account_id"`
}

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      *string   `gorm:"size:255"`
	Email     string    `gorm:"unique;size:255"`
	Password  string    `gorm:"size:255"`
	Image     *string   `gorm:"size:255"`
	Role      string    `gorm:"default:member;size:50"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
