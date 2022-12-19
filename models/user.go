package models

import (
	"gorm.io/gorm"
)

type User struct {

	// gorm.Model equals following (instead of writing all these 4, we can just use gorm.Model)
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`

	gorm.Model
	Email    string `gorm:"unique"` 	//make sure every email in our database is unique
	Password string
}
