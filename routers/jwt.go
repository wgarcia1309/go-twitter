package routers

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

var UserEmail string

var UserID string

func ProcessToken(tokenRaw string) (*models.Claim, bool, string, error) {
	jwtKey := []byte(os.Getenv("JWTKEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(tokenRaw, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tokenRaw = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tokenRaw, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return claims, false, string(""), err
	}

	if !token.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	_, found := db.EmailExist(claims.Email)
	if found {
		UserEmail = claims.Email
		UserID = claims.ID.Hex()
	}
	return claims, found, UserID, nil
}
