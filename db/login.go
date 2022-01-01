package db

import (
	"github.com/wgarcia1309/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(emailOrUsername, password string) (models.User, bool) {
	usr, exitsUsername, _ := UsernameExist(emailOrUsername)
	usr, existEmail, _ := EmailExist(emailOrUsername)
	if !existEmail || !exitsUsername {
		return usr, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usr, false
	}
	return usr, true
}
