package bd

import (
	"context"
	"time"

	"github.com/xavimg/TweetGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittorclone")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	//Del bson que te retorna result, extrae la ultima clave del campo insertado y saca el objectID
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
