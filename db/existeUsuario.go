package db

import (
	"context"
	"time"

	"github.com/Drakoxw/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*BuscarUsuario => busca si ya existe el usuario en la db*/
func BuscarUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoC.Database("Data-Disco")
	coll := db.Collection("Usuarios_t")

	condicion := bson.M{"email": email}

	var r models.Usuario

	err := coll.FindOne(ctx, condicion).Decode(&r)

	id := r.ID.Hex()
	if err != nil {
		return r, false, id
	}

	return r, true, id

}
