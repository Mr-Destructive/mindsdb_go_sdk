package main

import (
	"fmt"
	"os"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func main() {
	err := mindsdb.LoadEnvFromFile(".env")
	PanicError(err)
	email := os.Getenv("email")
	password := os.Getenv("password")
	api, err := connectors.Login(email, password)
	PanicError(err)

	server := mindsdb.Server{Api: api}
	params := map[string]string{"user": "abc", "password": "abc"}

	// Create a new project in mindsdb with server(cloud or localhost), name, engine, and additional params
	new_project := mindsdb.NewProject(&server, "test_sdk", "mindsdb", params)

	fmt.Println(new_project)
}
