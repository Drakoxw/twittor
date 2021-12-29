package jwt

import (
	"time"

	"github.com/Drakoxw/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerarToken => crea y firma el token */
func GenerarToken(u models.Usuario) (string, error) {

	clave := []byte("ErrorenlaFirma")

	payload := jwt.MapClaims{
		"email":    u.Email,
		"nombre":   u.Nombre,
		"apellido": u.Apellido,
		"_id":      u.ID.Hex(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tkStr, err := token.SignedString(clave)
	if err != nil {
		return tkStr, err
	}
	return tkStr, nil

}
