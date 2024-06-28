package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id uint `gorm:"primaryKey"`
	Username      string `gorm:"unique"`
	Password      string
	Name          string
	ContactPhone  string
	ContactEmail  string
	UserRole      string
}

