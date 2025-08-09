package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email        string `gorm:"uniqueIndex"`
    PasswordHash string
    RefreshToken string
}
