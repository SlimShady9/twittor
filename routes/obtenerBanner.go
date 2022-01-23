package routes

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/SlimShady9/twittor/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro del id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, _ := ioutil.ReadFile("uploads/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Error al copiar el Banner", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(OpenFile)))
	w.Write(OpenFile)

}
