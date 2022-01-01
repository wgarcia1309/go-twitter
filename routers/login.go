package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/jwt"
	"github.com/wgarcia1309/go-twitter/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application-json")
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, "Error decoding"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) == 0 {
		http.Error(rw, "Email de usuario requerido", http.StatusBadRequest)
		return
	}
	if len(u.Username) == 0 {
		http.Error(rw, "nombre de usuario requerido", http.StatusBadRequest)
		return
	}
	var model models.User
	var exist bool
	model, exist = db.Login(u.Email, u.Password)
	if !exist {
		model, exist = db.Login(u.Username, u.Password)
	}
	if !exist {
		http.Error(rw, "usuario y/o contraseña invalidos", http.StatusBadRequest)
		return
	}
	jwtkey, err := jwt.CreateJWT(model)
	if err != nil {
		http.Error(rw, "Error generating jwt"+err.Error(), http.StatusInternalServerError)

	}
	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
