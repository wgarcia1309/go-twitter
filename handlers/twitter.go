package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func CreateTweet(rw http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		http.Error(rw, "Error decoding :"+err.Error(), http.StatusBadRequest)
		return
	}

	tweet.UserID = UserID
	tweet.Date = time.Now()

	err = db.NewTweet(tweet)
	if err != nil {
		http.Error(rw, "Error saving tweet "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func GetTweets(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("id")
	if len(userID) < 1 {
		http.Error(w, "must send id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "must send page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page parameter must be greater than 0", http.StatusBadRequest)
		return
	}

	page64 := int64(page)
	response, ok := db.GetTweets(userID, page64)
	if !ok {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if response == nil {
		http.Error(w, "no tweets", http.StatusOK)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("id")
	if len(tweetID) < 1 {
		http.Error(w, "must send id parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(tweetID, UserID)
	if err != nil {
		http.Error(w, "something went wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
