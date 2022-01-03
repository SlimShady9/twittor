package bd

import (
	"context"
	"log"
	"time"

	"github.com/SlimShady9/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoTweets es la lectura de los tweets de un usuario en la BD */
func LeoTweets(ID string, pagina int64) ([]*models.DevulevoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("twittor")
	col := db.Collection("tweets")

	var resultado []*models.DevulevoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}
	for cursor.Next(context.TODO()) {
		var registro models.DevulevoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}
	return resultado, true
}
