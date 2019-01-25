package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type PasswordManager struct {
	Username          string `gorm:"unique; not null"`
	EncryptedPassword string `gorm:"not null" json:"-"`
}

func (u *PasswordManager) Encrypt(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(hashed)
	return nil
}

func (u *PasswordManager) VerifyPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}

func (u *PasswordManager) ValidatePassword(password string) (bool, error) {
	if len(password) == 0 {
		return false, errors.New("password can not be empty")
	} else if matched, _ := regexp.MatchString(`\s`, password); !matched {
		return false, errors.New("password can not contain any white space")
	}

	return true, nil
	// matched, _ := regexp.MatchString(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[a-zA-Z0-9\d@$!%*?&]{3,6}$`, password)
	// return matched
}
