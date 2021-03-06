package model

import (
	"github.com/google/uuid"
)

// File : Model with injected fields `ID`, `CreatedAt`, `UpdatedAt`
type File struct {
	Model

	Name string `gorm:"not null;`
	
	Tags string `gorm:"not null;`

	UserID uint // foreign key
	
	Uuid uuid.UUID // file UUID
}
