package main

import (
	"fmt"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := mindsdb.LoadEnvFromFile(".env")
	email := os.Getenv("email")
	password := os.Getenv("password")

	api, err := connectors.Login(email, password)
	//api, err = mindsdb.Connect("https://cloud.mindsdb.com", email, password)

	HandleError(err)
	logged_session := api.Session
	_, _, err = api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	HandleError(err)
	server := mindsdb.Server{Api: api}
	projects := server.ListProjects()
	for _, project := range projects {
		fmt.Println(project.Name)
	}
	project, err := server.GetProject("test_sdk_1")
	HandleError(err)
	query, _ := project.Query("SELECT * FROM test_sdk.models WHERE name='model_b';", "")
	project, err = server.GetProject("test_sdk")
	view := project.NewView("view_dj_ai", "SELECT response FROM test_sdk.model_b WHERE language='django' AND field='artificial intelligence';")
	fmt.Println(view)
	fmt.Println(query.ResultSet.Columns)
	db, _ := server.GetDatabase("test_sdk_1")
	res, _ := db.Query("SHOW FULL TABLES")
	fmt.Println(res.ResultSet.Columns)
	views, _ := project.ListViews()
	fmt.Println("VIEWS: ", views)
	//fmt.Println(db.Tables)
	//data, _ := db.Query("SELECT * FROM models WHERE name='model_dj'")
	//fmt.Println(data)
	view_index, _ := project.GetView("")
	fmt.Println(view_index)
	for _, table := range db.ListTables() {
		fmt.Println(table.Name)
	}
	/*
		params := map[string]string{"user": "abc", "password": "abc"}
		new_project := mindsdb_go_sdk.NewProject(&server, "test_sdk", "mindsdb", params)
		fmt.Println(new_project)
		project, err := server.GetProject("test_sdk_1")
		HandleError(err)
		fmt.Println(project)
		models := server.ListModels("test_sdk")
		fmt.Println(models)
		model, err := project.GetModel("model_dj")
		HandleError(err)
		fmt.Println(model)
		fmt.Println(server.DropProject("test_sdk_1"))
		project.DropModel("model_dj")

		prompt_tempate := ``
		model, err := project.NewModel("model_dj", "response", "openai", map[string]string{"model_name": "gpt-3.5-turbo", "prompt_template": prompt_tempate})
		HandleError(err)
		fmt.Println(model)
		resultSet := model.Predict("response", map[string]string{"question": "What should I name my custom fork of django?", "context": "creative"})
		fmt.Println(resultSet.Rows)
		server.GetModel("test_sdk_1", "model_dj")
		resultSet := server.Predict("test_sdk", "model_x", "response", map[string]string{"field": ""})
		fmt.Println(resultSet.Rows)
	*/
}
