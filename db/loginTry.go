package db

import (
	"fmt"
	"github.com/brayanzuritadev/citas/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginTry(email string, password string) (models.User, bool) {

	user, find := GetUser(email)
	if !find {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		fmt.Print(err.Error())
		return user, false
	}

	return user, true
}
