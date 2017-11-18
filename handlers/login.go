package handlers

import (
	"net/http"
	"github.com/partyzanex/test-task/model"
)

var (
	Auth map[string]string
	Work map[string]int32
)

func Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("pass")

	if Auth[login] == pass {
		w.WriteHeader(http.StatusOK)
	}

	user := &model.User{}
	err := user.Get(login, pass)
	if err == nil {
		Auth[login] = pass
		Work[login] = user.WorkNumber
	}

	w.WriteHeader(http.StatusBadRequest)
}

func ChangePass(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("pass")

	newPass := r.FormValue("newPass")

	if Auth[login] != pass {
		w.WriteHeader(http.StatusBadRequest)
	}

	user := &model.User{}
	user.Pass = newPass
	err := user.Save()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
