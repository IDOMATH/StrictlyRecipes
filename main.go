package main

import (
	"fmt"
	"github.com/IDOMATH/StrictlyRecipes/router"
	"net/http"
)

const port = "8080"

func main() {
	rtr := router.NewRouter()
	http.HandleFunc("/", rtr.Route)

	fmt.Println("Server running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
