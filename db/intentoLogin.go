package db

import (
	"github.com/Drakoxw/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin => buscar el usuario y compara las contraseñas*/
func IntentoLogin(email string, pass string) (models.Usuario, bool) {

	us, encontrado, _ := BuscarUsuario(email)
	if !encontrado {
		return us, false
	}

	passByte := []byte(pass)
	passDB := []byte(us.Pass)

	err := bcrypt.CompareHashAndPassword(passDB, passByte)
	if err != nil {
		return us, false
	}
	return us, true

}
