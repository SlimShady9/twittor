package routes

import (
	"net/http"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se encontro relacion para el usuario "+ID, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
