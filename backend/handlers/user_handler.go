package handlers

import "github.com/IDOMATH/StrictlyRecipes/db"

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}
