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
	project, err := server.GetProject("test_sdk")
	if err != nil {
		panic(err)
	}
	fmt.Println(project)

	// Create a new model in mindsdb with server(cloud or localhost), name, column to be predicted, engine, and additional params
	params := map[string]string{}
	model, err := project.NewModel("test_model", "something", "mindsdb", params)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}
