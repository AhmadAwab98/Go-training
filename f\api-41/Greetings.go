package main

import (
	"strings"
	"net/http"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// get the name from URL
	args := mux.Vars(r)
	name := args["name"]

	// respond to client
	w.Write([]byte("Hello, " + strings.Title(strings.ToLower(name))))
}


func main() {
	r := mux.NewRouter()

	// define endpoint and associate with handler
	r.HandleFunc("/hello/{name}", helloHandler).Methods("GET")

	// setup http server on port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}