package models

import "time"

type User struct {
	ID         uint       `gorm:"primaryKey"`
	Name       string     `gorm:"not null"`
	Email      string     `gorm:"unique;not null"`
	Phone      string     `gorm:"not null"`
	InviteCode string     `gorm:"unique;not null"`
	InvitedBy  *uint
	Points     int        `gorm:"default:1"`
	CreatedAt  time.Time
}

type RegisterRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	InviteCode string `json:"invite_code"`
}
