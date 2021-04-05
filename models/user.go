package Models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"not null"json:"name"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
