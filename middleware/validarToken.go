package middleware

import (
	"net/http"

	"github.com/Drakoxw/twittor/routers"
)

/*ValidarToken => validador de tokens*/
func ValidarToken(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			routers.ResponseErr(rw, "No Autorizado", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
