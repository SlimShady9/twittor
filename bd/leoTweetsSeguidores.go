package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/SlimShady9/twittor/models"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevulevoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevulevoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
