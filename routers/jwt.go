package routers

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

var UserEmail string

var UserID string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	jwtKey := []byte(os.Getenv("JWTKEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil {
		_, encontrado, _ := db.EmailExist(claims.Email)
		if encontrado {
			UserEmail = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, encontrado, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
