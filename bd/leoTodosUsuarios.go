package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SlimShady9/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTodosUsuarios(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	var encontrado, incluir bool
	for cur.Next(context.TODO()) {
		var elem models.Usuario
		err := cur.Decode(&elem)
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = elem.ID.Hex()
		incluir = false

		encontrado, _ = ConsultoRelacion(r)

		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}
		if r.UsuarioID == r.UsuarioRelacionID {
			incluir = false

		}
		if incluir {
			elem.Password = ""
			elem.Biografia = ""
			elem.SitioWeb = ""
			elem.Ubicacion = ""
			elem.Banner = ""
			elem.Email = ""

			fmt.Println("entro")
			results = append(results, &elem)
		}
	}
	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
