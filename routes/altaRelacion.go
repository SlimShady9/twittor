package routes

import (
	"net/http"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/models"
)

/* AltaRelacion realiza el registro de una relacion en la base de datos */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el alta de relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado dar de alta la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
