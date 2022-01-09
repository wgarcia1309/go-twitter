package handlers

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
	if len(u.Email) == 0 && len(u.Username) == 0 {
		http.Error(rw, "username or email required", http.StatusBadRequest)
		return
	}

	if len(u.Password) == 0 {
		http.Error(rw, "password required", http.StatusBadRequest)
		return
	}
	var model models.User
	model, ok := db.Login(u.Email, u.Username, u.Password)
	if !ok {
		http.Error(rw, "Wrong username or password", http.StatusBadRequest)
		return
	}
	jwt, err := jwt.CreateJWT(model)
	if err != nil {
		http.Error(rw, "Error creating jwt : "+err.Error(), http.StatusInternalServerError)
	}

	resp := models.RespuestaLogin{
		Token: jwt,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: expirationTime,
	})
}
