package bd

import (
	"github.com/SlimShady9/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentoLogin realza el intento de logueo de un usuario en la BD */
func IntentoLogin(email, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true

}
