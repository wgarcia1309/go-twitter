package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wgarcia1309/go-twitter/db"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "page parameter must be a integer greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.FindUsers(SessionUserID, pag, search, typeUser)
	if !status {
		http.Error(w, "error reading users", http.StatusBadRequest)
		return
	}
	if len(result) == 0 {
		http.Error(w, "no users found", http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
