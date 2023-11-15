package main

import (
	"strings"
	"net/http"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	name := args["name"]
	if name == "" {
		w.Write([]byte("No name given"))
		return
	}
	w.Write([]byte("Hello, " + strings.Title(strings.ToLower(name))))
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}