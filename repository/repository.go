package repository

import (
	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/router"
)

type Repository struct {
	Router *router.Router
	RH     *handlers.RecipeHandler
}

func NewRepository() *Repository {
	return &Repository{}
}
