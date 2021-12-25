package middlew

import (
	"net/http"

	"github.com/SlimShady9/twittor/bd"
)

/* ChequeoBD es la función que se encarga de verificar la conexión a la BD
* @returns true si la conexión es exitosa, false en caso contrario
 */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(rw, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
