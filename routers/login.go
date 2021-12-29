package routers

import (
	"encoding/json"
	"net/http"

	//"time"

	"github.com/Drakoxw/twittor/db"
	"github.com/Drakoxw/twittor/jwt"
	"github.com/Drakoxw/twittor/models"
)

/*Login => Logueador de agente*/
func Login(rw http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		ResponseErr(rw, "Usuario y/o Contraseña Erroneos", http.StatusBadRequest)
		return
	}

	if len(u.Email) == 0 {
		ResponseErr(rw, "Email es Requerido", http.StatusBadRequest)
		return
	}

	doc, existe := db.IntentoLogin(u.Email, u.Pass)
	if !existe {
		ResponseErr(rw, "Usuario y/o Contraseña Erroneos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerarToken(doc)
	if err != nil {
		ResponseErr(rw, "No se Pudo Generar el Token", http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(resp)

}
