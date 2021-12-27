package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/models"
)

/* Registro es la funcion para crear un nuevo usuario en la BD*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
