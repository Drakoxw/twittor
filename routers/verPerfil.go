package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Drakoxw/twittor/db"
)

func VerPerfil(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		ResponseErr(rw, "Se requiere un ID", 400)
		return
	}

	perfil, err := db.BuscarPerfil(id)
	if err != nil {
		ResponseErr(rw, "No se encontro perfil", 400)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(perfil)
}
