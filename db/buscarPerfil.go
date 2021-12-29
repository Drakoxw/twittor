package db

import (
	"context"
	"time"

	"github.com/Drakoxw/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscarPerfil => elimina el pass antes de volver el perfil*/
func BuscarPerfil(id string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoC.Database("Data-Disco")
	coll := db.Collection("Usuarios_t")

	var perfil models.Usuario

	objId, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{
		"_id": objId,
	}

	err := coll.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Pass = ""
	if err != nil {
		return perfil, err
	}
	return perfil, nil

}
