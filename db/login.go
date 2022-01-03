package db

import (
	"github.com/wgarcia1309/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, username, password string) (models.User, bool) {
	var usr models.User
	var exist bool
	if username != "" {
		usr, exist, _ = UsernameExist(username)
	}
	if email != "" {
		usr, exist, _ = EmailExist(email)
	}
	if !exist {
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
