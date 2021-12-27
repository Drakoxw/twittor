package db

import "golang.org/x/crypto/bcrypt"

/*Encriptar => encripta culaquier string */
func Encriptar(x string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(x), 6)
	return string(bytes), err

}
