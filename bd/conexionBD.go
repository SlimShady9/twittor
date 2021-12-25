package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = conectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://jondo:Juancho9@cluster0.y6kuj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* conectarBD para conectarse a la BD
* @returns cliente de la BD
 */
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Conexion exitosa a la BD")
	return client
}

/* checkConnection es la función que se encarga de verificar la conexión a la BD
* @returns true si la conexión es exitosa, false en caso contrario
 */
func CheckConnection() bool {
	err := MongoC.Ping(context.TODO(), nil)
	return err == nil
}
