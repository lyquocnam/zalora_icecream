package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int    `gorm:"primary_key" json:"id"`
	Username       string `gorm:"unique; not null"`
	DisplayName    string `gorm:"not null" json:"display_name"`
	HashedPassword string `gorm:"not null" json:"-"`
}

func (u *User) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(hashed)
	return nil
}

func (u *User) ComparePassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
