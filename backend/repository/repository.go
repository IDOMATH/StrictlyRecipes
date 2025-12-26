package repository

import (
	"net/http"

	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/render"
	"github.com/IDOMATH/StrictlyRecipes/types"
)

type Repository struct {
	RH *handlers.RecipeHandler
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.go.html", &types.TemplateData{PageTitle: "Home"})
}
