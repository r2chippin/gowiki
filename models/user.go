package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Nickname     string
	Email        string `gorm:"unique;not null"`
	PasswordHash []byte `gorm:"not null"`
}
