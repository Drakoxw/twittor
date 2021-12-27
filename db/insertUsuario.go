package db

import (
	"context"
	"time"

	"github.com/Drakoxw/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarUser => inserta el usuario */
func InsertarUser(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoC.Database("Data-Disco")
	coll := db.Collection("Usuarios_t")

	result, err := coll.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}
