package middleware

import (
	"net/http"

	"github.com/wgarcia1309/go-twitter/routers"
)

func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "token error : "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
