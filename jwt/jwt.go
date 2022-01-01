package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wgarcia1309/go-twitter/models"
)

func CreateJWT(t models.User) (string, error) {

	miClave := []byte(os.Getenv("JWTKEY"))

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Name,
		"apellidos":        t.Lastname,
		"fecha_nacimiento": t.Birthdate,
		"biografia":        t.Bio,
		"ubicacion":        t.Location,
		"sitioweb":         t.Website,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
