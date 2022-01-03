package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SlimShady9/twittor/bd"
	"github.com/SlimShady9/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el registro, reintente nuevamente", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
