package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wgarcia1309/go-twitter/db"
)

func GetFeed(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "must send page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page parameter must be a integer greater than 0", http.StatusBadRequest)
		return
	}

	response, ok := db.GetFeed(SessionUserID, page)
	if !ok {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
