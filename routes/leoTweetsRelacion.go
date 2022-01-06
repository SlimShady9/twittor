package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SlimShady9/twittor/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, status := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
