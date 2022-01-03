package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SlimShady9/twittor/middlew"
	"github.com/SlimShady9/twittor/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores creo el reouter y configuro CORS	*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT((routes.VerPerfil)))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT((routes.ModificarPerfil)))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT((routes.GraboTweet)))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routes.LeoTweets))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
