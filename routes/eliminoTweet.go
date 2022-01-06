package routes

import (
	"net/http"

	"github.com/SlimShady9/twittor/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	error := bd.BorroTweet(ID, IDUsuario)
	if error != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+error.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
