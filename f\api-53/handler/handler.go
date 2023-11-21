package handler

import (
	"encoding/json"
	"go-training/csv"
	"net/http"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("Awwab")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

 return tokenString, nil
}

var users = map[string]string{
	"Awwab98": "KhanAwab",
	"Ahmad89": "Ahmad1989",
}

// represent JSON structure in response body

type bodyResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Time string `json:"time"`
}

// represent expected JSON structure in request body
type bodyRequest struct {
	CSVFilepath string `json:"CSVFilepath"`
}

type credentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type userToken struct {
	Username string `json:"Username"`
	Token string `json:"token"`
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var cred credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		return
	}

	// get the password from our user map
	password, ok := users[cred.Username]

	if !ok || password != cred.Password {
		w.Write([]byte("Unauthorised"))
		return
	}

	// declare the token with the algorithm used for signing
	token := jwt.New(jwt.SigningMethodHS256)

	// create the JWT string
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		// f error creating JWT write internal server error
		w.Write([]byte("Internal Server Error"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path: "/CSV",
	})
}

func GuestHandler(w http.ResponseWriter, r *http.Request) {

	// prepare response
	response := bodyResponse{
		Message : "Hello, " + "Guest",
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

func CSVHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("UserToken")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.Write([]byte("Unauthorized"))
			return
		}
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	tkn, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.Write([]byte("Unauthorized"))
			return
		}
		return
	}
	if !tkn.Valid {
		w.Write([]byte("Unauthorized"))
		return
	}

	// decode JSON into bodyRequest
	decoder := json.NewDecoder(r.Body)
	var request bodyRequest
	err = decoder.Decode(&request)
	if err != nil {
        return
    }

	JSONresponse := csv.ParseCSV(request.CSVFilepath)

	// respond to client
	w.Write([]byte(JSONresponse))
}