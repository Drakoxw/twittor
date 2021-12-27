package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://DrakoMaster:wolfW12345@drako-db.fguhd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD => Conexion a la Base de Datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la Base de Datos")
	return client
}

/*CheckConect() => pin a la base de datos  */
func CheckConect() int {
	err := MongoC.Ping(ctx, nil)
	if err != nil {
		return 0
	}
	return 1
}

/* eee
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
} */
