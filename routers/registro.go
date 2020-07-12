package routers

import (
	"encoding/json"
	"github.com/valianx/clon-twitter/bd"
	"github.com/valianx/clon-twitter/models"
	"net/http"
)

//Registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese Email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al ingresar el usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Ocurrio un error al ingresar el usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
