package main

import (
	"fmt"
	"os"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// read email and password from the .env file
	err := mindsdb.LoadEnvFromFile(".env")
	PanicError(err)
	email := os.Getenv("email")
	password := os.Getenv("password")

	// Login in with an email and password
	api, err := connectors.Login(email, password)
	PanicError(err)

	logged_session := api.Session
	// Query all the model names from the mindsdb cloud
	rows, cols, err := api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	PanicError(err)

	fmt.Println(rows, cols)
}
