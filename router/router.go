package router

import (
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
		// TODO: Handle 404
		// TODO: I"m too drunk to actually make smething thonight.  Just thso once

	}
}
