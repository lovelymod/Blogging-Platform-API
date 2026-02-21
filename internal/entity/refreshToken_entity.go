package entity

import (
	"time"
)

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"userID" gorm:"not null;index"`
	Token     string    `json:"token" gorm:"not null;unique"`
	Jti       string    `json:"jti" gorm:"not null;unique"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	IsRevoked bool      `json:"isRevoked" gorm:"default:false"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
