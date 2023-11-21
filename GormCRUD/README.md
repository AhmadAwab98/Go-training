# Go training 
Repository for the training of Go language.

## Description

The project contains:
- use go-chi/chi to add route
- use "gorm.io/driver/pq" and "gorm.io/gorm" for database.
- add a handler function to
    - create owner
    - delete owner
    - update owner
    - read all owners or by id
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
 go run main.go
```

#### Select all
- On the Postman set the method to Get, write localhost:8080/object in URL field.
- Click send, you will see all users present in the database.

#### Select by id
- On the Postman set the method to Get, write localhost:8080/object/id in URL field.
- Click send, you will see the user with the required id present in the database.


#### Create
- On the Postman set the method to Post, write localhost:8080/object in URL field and in the body write
```
{
    "Name": "Awab"
    "Email": "awwab@gmail.com"
}
```
- Click send, you will see created user.


#### Update
- On the Postman set the method to Patch, write localhost:8080/object/id in URL field and in the body write
```
{
    "Name": "Awab"
    "Email": "awwab@gmail.com"
}
```
- Click send, you will see updated.


#### Delete
- On the Postman set the method to Delete, write localhost:8080/object/id in URL field.
- Click send, you will see deleted.


## Authors

Ahmad Awab
