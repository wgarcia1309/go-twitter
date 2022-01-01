package routers

import (
	"encoding/json"
	"net/http"

	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/models"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Error decoding"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(rw, "Email de usuario requerido", http.StatusBadRequest)
		return
	}
	if len(t.Username) == 0 {
		http.Error(rw, "nombre de usuario requerido", http.StatusBadRequest)
		return

	}
	if len(t.Password) <= 6 {
		http.Error(rw, "contraseña de usuario requerido", http.StatusBadRequest)
		return
	}
	_, founded, _ := db.EmailExist(t.Email)
	if founded {
		http.Error(rw, "ya existe un usuario regristrado con ese email", http.StatusBadRequest)
		return
	}
	_, founded, _ = db.UsernameExist(t.Username)
	if founded {
		http.Error(rw, "ya existe un usuario regristrado con ese nombre de usuario", http.StatusBadRequest)
		return
	}

	_, status, err := db.NewUser(t)
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
