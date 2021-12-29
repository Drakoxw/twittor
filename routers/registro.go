package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Drakoxw/twittor/db"
	"github.com/Drakoxw/twittor/models"
)

type HttpErr2 struct {
	Err bool   `json:"err"`
	Msg string `json:"msg"`
}

/*Registro => registra el usuario */
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		ResponseErr(rw, "No hay Data", http.StatusNotAcceptable)
		return
	}
	if len(t.Email) == 0 {
		ResponseErr(rw, "Email es requerido", http.StatusNoContent)
		return
	}
	if len(t.Pass) < 6 {
		ResponseErr(rw, "Contaseña muy corta", http.StatusNoContent)
		return
	}

	us, encontrado, _ := db.BuscarUsuario(t.Email)
	if encontrado {
		ResponseErr(rw, "Usuario ya Existe", 409)
		return
	}

	t.Pass, err = db.Encriptar(t.Pass)
	if err != nil {
		ResponseErr(rw, "UNo se pudo Encriptar la Contraseña", http.StatusNotAcceptable)
		return
	}

	_, status, err := db.InsertarUser(t)

	if err != nil {
		ResponseErr(rw, "Error en insert User", 404)
		return
	}

	if !status {
		ResponseErr(rw, "Error2 en insert User", 404)
		return
	}

	rw.WriteHeader(201)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(us)

}
