package routers

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func Register(rw http.ResponseWriter, r *http.Request) {
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
