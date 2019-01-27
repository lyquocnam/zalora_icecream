package services

import (
	"github.com/lyquocnam/zalora_icecream/lib"
	"github.com/lyquocnam/zalora_icecream/models"
)

func FindUserByUsername(username string) *models.User {
	user := &models.User{}
	lib.DB.First(&user, "username = ?", username)
	return user
}
