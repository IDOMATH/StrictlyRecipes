package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}
