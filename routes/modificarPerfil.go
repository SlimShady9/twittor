package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+err.Error(), 400)
		return
	}
	if status {
		http.Error(w, "No se modifico el registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
