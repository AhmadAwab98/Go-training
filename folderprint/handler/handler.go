package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// represent expected JSON structure in request body
type bodyRequest struct {
	Path string `json:"path"`
}

// represent JSON structure of response body
type bodyResponse struct {
	Folders []string `json:"folders"`
	Files []string `json:"files"`
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	// decode JSON into bodyRequest
	decoder := json.NewDecoder(r.Body)
	var request bodyRequest
	err:= decoder.Decode(&request)
	if err != nil {
        return
    }

	var response bodyResponse

	// prepare response
	recPrepareResponse(request.Path, &response)

	if err != nil {
        return
    }

	// converting to JSON
	JSONresponse, errjson := json.MarshalIndent(response,"","	")
	if errjson != nil {
		return
	}

	// respond to client
	w.Write([]byte(JSONresponse))
}

func recPrepareResponse(Path string, response *bodyResponse) {
	info, err := os.Stat(Path)
	
	// getting the name of the directory/file
	name := strings.Split(Path, "/")
	switch {
	// if broken or not opening return 
	case err != nil:
		return
	case info.IsDir():
	// add the folder in folders array of response 
		response.Folders = append(response.Folders, name[len(name) - 1])
		files, err := os.ReadDir(Path)
		if err != nil {
			return
		}
		for _, file := range files {

			// call the function recursively on every directory
			recPrepareResponse(filepath.Join(Path, file.Name()), response)
		}
	default:
		// add the files in files array of response 
		response.Files = append(response.Files, name[len(name) - 1])
		return
	}
}