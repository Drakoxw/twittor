package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Drakoxw/twittor/db"
	"github.com/Drakoxw/twittor/models"
)

/*Registro => registra el usuario */
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Error de Usuario: "+err.Error(), http.StatusNotAcceptable)
		return
	}
	if len(t.Email) == 0 {
		http.Error(rw, "Email es requerido: "+err.Error(), http.StatusNoContent)
		return
	}
	if len(t.Pass) < 6 {
		http.Error(rw, "Contaseña muy corta: "+err.Error(), http.StatusNoContent)
		return
	}

	_, encontrado, _ := db.BuscarUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "Usuario ya Existe: "+err.Error(), http.StatusConflict)
		return
	}

	t.Pass, err = db.Encriptar(t.Pass)
	if err != nil {
		http.Error(rw, "Error en Passw: "+err.Error(), http.StatusNotAcceptable)
		return
	}

	_, status, err := db.InsertarUser(t)

	if err != nil {
		http.Error(rw, "Error en insert User: "+err.Error(), 404)
		return
	}

	if !status {
		http.Error(rw, "Error en insert User: "+err.Error(), 404)
		return
	}

	rw.WriteHeader(201)

}
