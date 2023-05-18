package main

import (
	"fmt"
	"log"
	"mindsdb_go_sdk"
	"mindsdb_go_sdk/connectors"
	"net/http"
	"net/url"
	"time"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	api := connectors.RestAPI{}
	apiUrl, err := url.Parse("https://cloud.mindsdb.com")
	api.Url = apiUrl
	api.Email = "abc@mail.com"
	api.Password = "**Password**"
	session := &http.Client{
		Timeout: time.Second * 10,
	}
	api.Session = session
	req, err := api.Login()
	HandleError(err)
	logged_session := req.Session
	_, _, err = api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	HandleError(err)
	server := mindsdb_go_sdk.Server{}
	server = *server.NewServer(&api)
	projects := server.ListProjects()
	for _, project := range projects {
		fmt.Println(project.Name)
	}
	/*
		params := map[string]string{"user": "abc", "password": "abc"}
		new_project := mindsdb_go_sdk.NewProject(&server, "test_sdk", "mindsdb", params)
		fmt.Println(new_project)
		project, err := server.GetProject("test_sdk")
		HandleError(err)
		fmt.Println(project)
		models := server.ListModels("test_sdk")
		fmt.Println(models)
		model := server.GetModel("test_sdk_1", "model_x")
		fmt.Println(model)

			fmt.Println(server.DropProject("test_sdk_1"))
			prompt_tempate := `Give some project ideas for {{field}} in the domain of {{area}}
			with programmin language as {{language}}`
			model := server.NewModel("test_sdk_1", "model_x", "response", "openai", map[string]string{"model_name": "gpt-3.5-turbo", "prompt_template": prompt_tempate})
			fmt.Println(model)
			resultSet := server.Predict("test_sdk_1", "model_x", "response", map[string]string{"field": "backend", "area": "space", "language": "golang"})
			fmt.Println(resultSet.Rows)
			server.GetModel("test_sdk_1", "model_x")
			resultSet := server.Predict("test_sdk", "model_x", "response", map[string]string{"field": ""})
			fmt.Println(resultSet.Rows)
	*/
}
