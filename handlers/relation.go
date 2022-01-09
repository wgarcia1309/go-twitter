package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func Follow(w http.ResponseWriter, r *http.Request) {

	FollowerID := r.URL.Query().Get("id")
	if len(FollowerID) < 1 {
		http.Error(w, "must send id parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = SessionUserID
	t.UserRelationID = FollowerID

	err := db.Follow(t)
	if err != nil {
		http.Error(w, "something when wrong "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	FollowerID := r.URL.Query().Get("id")
	if len(FollowerID) < 1 {
		http.Error(w, "must send id parameter", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = SessionUserID
	t.UserRelationID = FollowerID

	err := db.Unfollow(t)
	if err != nil {
		http.Error(w, "something when wrong "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetRelation(w http.ResponseWriter, r *http.Request) {
	FollowerID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = SessionUserID
	t.UserRelationID = FollowerID

	var response models.RelationRetrieved

	err := db.GetRelation(t)
	if err != nil {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
