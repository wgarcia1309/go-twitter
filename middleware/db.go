package middleware

import (
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !db.CheckConecction() {
			http.Error(rw, "can't connect with db", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
