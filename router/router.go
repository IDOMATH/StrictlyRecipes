package router

import (
	"github.com/IDOMATH/StrictlyRecipes/render"
	"github.com/IDOMATH/StrictlyRecipes/types"
	"net/http"
	"strings"
)

type Router struct {
	urlIndex int
}

func (rtr *Router) Route(w http.ResponseWriter, r *http.Request) {
	rtr.urlIndex = 1
	url := strings.Split(r.URL.Path, "/")
	switch url[rtr.urlIndex] {
	case "":
		// TODO: Handle home
	default:
		render.Template(w, r, "error-404.go.html", &types.TemplateData{PageTitle: "Not Found"})
	}
}
