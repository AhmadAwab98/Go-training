package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/redis/go-redis/v9"
	"context"
	"time"
)
var ctx = context.Background()
var rdb *redis.Client
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", Password: "", DB: 0,
	})
}

// represent expected JSON structure in request body
type bodyRequest struct {
	Path string `json:"path"`
}

// represent JSON structure of response body
type bodyResponse struct {
	Folders []string `json:"folders" redis:"folders"`
	Files []string `json:"files" redis:"files"`
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	defer rdb.Close()
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

	// converting to JSON
	JSONresponse, errjson := json.Marshal(response)
	if errjson != nil {
		return
	}

	// caching response for 5 minutes
	err = rdb.HSet(ctx, "responseCache", "cachedData", JSONresponse).Err()
	err = rdb.Expire(ctx, "responseCache", 5*time.Minute).Err()

	// get the response from cache and convert it into json format
	responsecached, _ := rdb.HGetAll(ctx,"responseCache").Result()
	err = json.Unmarshal([]byte(responsecached["cachedData"]), &response)
	JSONresponse, errjson = json.MarshalIndent(response, "", "	")

	// respond to client
	w.Write([]byte(JSONresponse))
}

func recPrepareResponse(Path string, response *bodyResponse) {
	info, err := os.Stat(Path)

	// getting the name of the directory/file
	name := strings.Split(Path, "/")
	switch {
	case err != nil:
		// if broken or not opening return
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