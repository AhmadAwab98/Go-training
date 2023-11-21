# Go training 
Repository for the training of Go language.

## Description

The project contains:
- use go-chi/chi to add route
- add a handler function to
    - convert CSV to JSON case
    - display JSON on web

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
 go run CSVtoJSON.go
```

- On the Postman set the method to Get/Post, write localhost:8080/hello in URL field and in the body write
```
{
    "CSVFilepath": "data.csv"
}
```
- Click send, you will see the result in the body tag of the output.

## Authors

Ahmad Awab
