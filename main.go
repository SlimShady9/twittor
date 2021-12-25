package main

import (
	"log"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/handlers"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
