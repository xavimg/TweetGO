package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncriptarPassword para enncriptar la contraseña*/

func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
