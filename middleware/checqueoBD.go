package middleware

import (
	"net/http"

	"github.com/Drakoxw/twittor/db"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConect() == 0 {
			http.Error(rw, "Error n conexión", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
