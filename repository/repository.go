package repository

import (
	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/render"
	"github.com/IDOMATH/StrictlyRecipes/types"
	"net/http"
	"regexp"
	"strings"
)

type Repository struct {
	urlIndex int
	RH       *handlers.RecipeHandler
}

func NewRepository() *Repository {
	return &Repository{}
}

var regexNumber = regexp.MustCompile(`\d`)

func (repo *Repository) Route(w http.ResponseWriter, r *http.Request) {
	repo.urlIndex = 1
	url := strings.Split(r.URL.Path, "/")
	switch url[repo.urlIndex] {
	case "":
		handleHome(w, r)
	case "recipes":
		repo.routeRecipes(w, r)
	default:
		render.Template(w, r, "error-404.go.html", &types.TemplateData{PageTitle: "Not Found"})
	}
}

func (repo *Repository) routeRecipes(w http.ResponseWriter, r *http.Request) {
	repo.urlIndex++
	url := strings.Split(r.URL.Path, "/")
	switch {
	case regexNumber.MatchString(url[repo.urlIndex]):
		repo.RH.HandleGetRecipeById(w, r, url[repo.urlIndex])
		return
	}
	repo.RH.HandleGetAllRecipes(w, r)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.go.html", &types.TemplateData{PageTitle: "Home"})
}
