package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC es el objeto de conexion a la BD*/
var MongoC, err = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://xavi:v6vpxdkd@twitterclone.bmk2f.mongodb.net/twittorclone?retryWrites=true&w=majority")

/* ConectarBD Es la funcion que nos permite crear la BD */
func ConectarBD() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	log.Println("Conexion exitosa con la BD")

	return client, nil
}

/*ChequeoConn es el Ping a la BD */

func ChequeoConn() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1

}
