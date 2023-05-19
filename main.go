package main

import (
	"fmt"
	"log"
	"mindsdb"
	"mindsdb/connectors"
	"net/http"
	"net/url"
	"os"
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
	err = mindsdb.LoadEnvFromFile(".env")
	api.Email = os.Getenv("email")
	api.Password = os.Getenv("password")
	session := &http.Client{
		Timeout: time.Second * 10,
	}
	api.Session = session
	req, err := api.Login()
	HandleError(err)
	logged_session := req.Session
	_, _, err = api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	HandleError(err)
	server := mindsdb.Server{}
	server = *server.NewServer(&api)
	projects := server.ListProjects()
	for _, project := range projects {
		fmt.Println(project.Name)
	}
	/*
		params := map[string]string{"user": "abc", "password": "abc"}
		new_project := mindsdb_go_sdk.NewProject(&server, "test_sdk", "mindsdb", params)
		fmt.Println(new_project)
	*/
	project, err := server.GetProject("test_sdk_1")
	HandleError(err)
	fmt.Println(project)
	//models := server.ListModels("test_sdk")
	//fmt.Println(models)
	model := project.GetModel("model_dj")
	//fmt.Println(model)
	//fmt.Println(server.DropProject("test_sdk_1"))
	//project.DropModel("model_dj")

	/*
		prompt_tempate := ``
		model, err := project.NewModel("model_dj", "response", "openai", map[string]string{"model_name": "gpt-3.5-turbo", "prompt_template": prompt_tempate})
		HandleError(err)
		fmt.Println(model)
	*/
	resultSet := model.Predict("response", map[string]string{"question": "What should I name my custom fork of django?", "context": "creative"})
	fmt.Println(resultSet.Rows)
	//server.GetModel("test_sdk_1", "model_dj")
	//resultSet := server.Predict("test_sdk", "model_x", "response", map[string]string{"field": ""})
	//fmt.Println(resultSet.Rows)
}
