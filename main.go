package main

import (
	"log"

	"github.com/Drakoxw/twittor/db"
	"github.com/Drakoxw/twittor/handlers"
)

// heroku: https://drako-twittor.herokuapp.com/
func main() {
	if db.CheckConect() == 0 {
		log.Fatal("Sin connexión a la Base de Datos")
		return
	}

	handlers.Manejadores()

}
