package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xavimg/TweetGO/bd"
	"github.com/xavimg/TweetGO/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es obligatorio", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password es menor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario con este email", 400)
		return

	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el registro de usuario", 400)
		return
	}

}
