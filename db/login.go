package db

import (
	"github.com/wgarcia1309/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, username, password string) (models.User, bool) {
	var usr models.User
	var found bool
	if username != "" {
		usr, found = UsernameExist(username)
	}
	if email != "" {
		usr, found = EmailExist(email)
	}
	if !found {
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
