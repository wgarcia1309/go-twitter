package routers

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func NewUser(rw http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, "Error decoding"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) == 0 {
		http.Error(rw, "Email de usuario requerido", http.StatusBadRequest)
		return
	}
	if len(u.Username) == 0 {
		http.Error(rw, "nombre de usuario requerido", http.StatusBadRequest)
		return

	}
	if len(u.Password) < 6 {
		http.Error(rw, "contraseña de usuario requerido", http.StatusBadRequest)
		return
	}
	_, founded, _ := db.EmailExist(u.Email)
	if founded {
		http.Error(rw, "ya existe un usuario regristrado con ese email", http.StatusBadRequest)
		return
	}
	_, founded, _ = db.UsernameExist(u.Username)
	if founded {
		http.Error(rw, "ya existe un usuario regristrado con ese nombre de usuario", http.StatusBadRequest)
		return
	}

	_, status, err := db.NewUser(u)
	if err != nil {
		http.Error(rw, "error saving in db"+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "can't save user in db", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.GetUserProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = db.UpdateUser(t, UserID)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
