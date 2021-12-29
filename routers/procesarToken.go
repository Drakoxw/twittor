package routers

import (
	"errors"
	"strings"

	"github.com/Drakoxw/twittor/db"
	"github.com/Drakoxw/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email => del usuario*/
var Email string

/*IdUser => del user*/
var IdUser string

/*ProcesarToken => verifica el token*/
func ProcesarToken(token string) (*models.Claim, bool, string, error) {

	clave := []byte("ErrorenlaFirma")
	claim := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, "", errors.New("formato de token no valido")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return clave, nil
	})
	if err == nil {
		_, encontrado, Id := db.BuscarUsuario(claim.Email)
		if encontrado {
			Email = claim.Email
			IdUser = Id
		}
		return claim, encontrado, Id, nil
	}

	if !tkn.Valid {
		return claim, false, string(""), errors.New("token vencido")
	}
	return claim, false, "", err

}
