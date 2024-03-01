package types

import "github.com/IDOMATH/StrictlyRecipes/handlers"

type Repository struct {
	RH *handlers.RecipeHandler
}

func NewRepository() *Repository {
	return &Repository{}
}
