package main

import (
	"fmt"
	"github.com/IDOMATH/StrictlyRecipes/router"
	"github.com/IDOMATH/StrictlyRecipes/util"
	"net/http"
)

func main() {
	port := util.EnvOrDefault("PORT", "8080")

	rtr := router.NewRouter()
	http.HandleFunc("/", rtr.Route)

	fmt.Println("Server running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
