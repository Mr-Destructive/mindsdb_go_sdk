package main

import (
	"fmt"
	"os"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func main() {
	err := mindsdb.LoadEnvFromFile(".env")
	mindsdb.HandleError(err)
	email := os.Getenv("email")
	password := os.Getenv("password")
	api, err := connectors.Login(email, password)
	mindsdb.HandleError(err)

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

	// Get a model with name inside a project
	model, err = project.GetModel("model_a")
	fmt.Println(model.Status)
	mindsdb.HandleError(err)

	// Get status of a model using the object
	model_status := model.Status
	fmt.Println(model_status)

	// Get status of a model using the object(reloaded sql query)
	status := model.GetStatus()
	fmt.Println(status)

	// Retrain model if the new data is available
	retrained_model, err := model.Retrain()
	if retrained_model != nil {
		fmt.Println(retrained_model)
		// model = retrained_model
	}
	fmt.Println(model, err)

	// FineTuned Model
	fmt.Println(model)
	finetuned_model, err := model.FineTune("test_sdk", "SELECT * FROM test_sdk.model_b WHERE language = 'go'")
	fmt.Println(finetuned_model, err)
}
