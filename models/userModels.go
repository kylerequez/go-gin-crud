package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string         `json:"name"`
	Email     *string        `gorm:"uniqueIndex" json:"email"`
	Age       *uint8         `json:"age"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
