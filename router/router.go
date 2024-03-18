package router

import (
	"github.com/IDOMATH/StrictlyRecipes/render"
	"github.com/IDOMATH/StrictlyRecipes/types"
	"net/http"
	"regexp"
	"strings"
)

var regexNumber = regexp.MustCompile(`\d`)

type Router struct {
	urlIndex int
}

func NewRouter() *Router {
	return &Router{}
}

func (rtr *Router) Route(w http.ResponseWriter, r *http.Request) {
	rtr.urlIndex = 1
	url := strings.Split(r.URL.Path, "/")
	switch url[rtr.urlIndex] {
	case "":
		handleHome(w, r)
	case "recipes":
		rtr.routeRecipes(w, r)
	default:
		render.Template(w, r, "error-404.go.html", &types.TemplateData{PageTitle: "Not Found"})
	}
}

func (rtr *Router) routeRecipes(w http.ResponseWriter, r *http.Request) {
	rtr.urlIndex++
	url := strings.Split(r.URL.Path, "/")
	switch {
	case regexNumber.MatchString(url[rtr.urlIndex]):
		// TODO: handle get recipe by ID and do the regex on case
	}
	render.Template(w, r, "all-recipes.go.html", &types.TemplateData{PageTitle: "All Recipes"})

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.go.html", &types.TemplateData{PageTitle: "Home"})
}
