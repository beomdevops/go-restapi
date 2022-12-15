package models

import "gorm.io/gorm"

type User struct {
	ID   int `gorm:"primaryKey; AUTO_INCREMENT"`
	Name string
	gorm.Model
}

func NewUser(name string) *User {
	return &User{Name: name}
}
