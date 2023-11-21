package handler

import (
	"encoding/json"
	"go-training/csv"
	"net/http"
)

// represent expected JSON structure in request body
type bodyRequest struct {
	CSVFilepath string `json:"CSVFilepath"`
}

func CSVHandler(w http.ResponseWriter, r *http.Request) {
	// decode JSON into bodyRequest
	decoder := json.NewDecoder(r.Body)
	var request bodyRequest
	err:= decoder.Decode(&request)
	if err != nil {
        return
    }

	JSONresponse := csv.ParseCSV(request.CSVFilepath)

	// respond to client
	w.Write([]byte(JSONresponse))
}