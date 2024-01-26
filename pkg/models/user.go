package models

import "gorm.io/gorm"

type UserDetail struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:32"`
	PasswordHash string
	Name         string
	Email        string
	Address      string
}
