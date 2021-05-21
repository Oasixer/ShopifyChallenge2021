package model

import (
	// "github.com/jinzhu/gorm"
)

// User : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type User struct {
	Model
	Email     string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"not null"`
	Username string `gorm:"type:varchar(50);not null"`
  Files []File
}
