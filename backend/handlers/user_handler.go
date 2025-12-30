package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IDOMATH/StrictlyRecipes/db"
	"github.com/IDOMATH/StrictlyRecipes/types"
)

type UserHandler struct {
	userStore *db.UserStore
}

func NewUserHandler(userStore *db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (*UserHandler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// log.Error("HandlePostUser", "error decoding user to json from body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (*UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// log.Error("HandleLogin", "error decoding user to json from body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
