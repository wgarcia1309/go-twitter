package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func NewUser(rw http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, "Error decoding :"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) == 0 {
		http.Error(rw, "email required", http.StatusBadRequest)
		return
	}
	if len(u.Username) == 0 {
		http.Error(rw, "username required", http.StatusBadRequest)
		return

	}
	if len(u.Password) < 6 {
		http.Error(rw, "password required", http.StatusBadRequest)
		return
	}
	_, found := db.EmailExist(u.Email)
	if found {
		http.Error(rw, "email already in use", http.StatusBadRequest)
		return
	}
	_, found = db.UsernameExist(u.Username)
	if found {
		http.Error(rw, "username already in use", http.StatusBadRequest)
		return
	}

	_, err = db.NewUser(u)

	if err != nil {
		http.Error(rw, "error saving in db :"+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

func GetProfile(rw http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("id")
	if len(userID) < 1 {
		http.Error(rw, "must send ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.GetUserProfile(userID)
	if err != nil {
		http.Error(rw, "something went wrong "+err.Error(), http.StatusNotFound)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "wrong information"+err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateUser(t, UserID)
	if err != nil {
		http.Error(w, "something went wrong"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
