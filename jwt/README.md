# Go training 
Repository for the training of Go language.

## Description

The project contains:
- use go-chi/chi to add route
- use golang-jwt/jwt to add authentication
- add a handler functions to convert CSV to JSON and didplay JSON if authenticated
- add a handler functions to display hello guest if not authenticated 

## Getting Started

### Language

Git

Go

### Execution

Continued commands from main readme.md

- Run the go mod tidy command to add any missing dependencies to the go.mod file
```
go mod tidy
```

- Run the go file
```
 go run jwt.go
```

- On the Postman set the method to Get, write localhost:8080/auth/signin in URL field and in the body write
```
{
    "Username": "Awwab9",
    "Password": "KhanAwab"
}
```

- Click send, if not authenticated you will get Unauthorized

- After that on the Postman set the method to Get, write localhost:8080/CSV in URL field and in the body write
```
{
    "CSVFilepath": "data.csv"
}
```

- Click send, if authenticated you will see the result in the body tag of the output
  
- If not authorized you can visit localhost:8080/guest. You will see a JSON body response of Hello, guest and time

## Authors

Ahmad Awab

