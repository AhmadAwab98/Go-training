package main

import (
	"encoding/json"
	"strings"
	"net/http"
	"github.com/go-chi/chi"
	"time"
)

// represent expected JSON structure in request body
type bodyRequest struct {
	Name string `json:"name"`
}

// represent JSON structure of response body
type bodyResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Time string `json:"time"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// decode JSON into bodyRequest
	decoder := json.NewDecoder(r.Body)
	var request bodyRequest
	err:= decoder.Decode(&request)
	if err != nil {
        return
    }

	// prepare response
	response := bodyResponse{
		Message : "Hello, " + strings.Title(strings.ToLower(request.Name)),
		Code : 200,
		Time : time.Now().Format("2006-01-02 15:04:05")}
	
	// converting to JSON
	JSONresponse, errjson := json.MarshalIndent(response,"","	")
	if errjson != nil {
		return
	}

	// respond to client
	w.Write([]byte(JSONresponse))
}

func main() {
	r := chi.NewRouter()

	// define endpoint and associate with handler
	r.Get("/hello", helloHandler)

	// setup http server on port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
