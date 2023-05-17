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

func Handle_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	api := connectors.RestAPI{}
	apiUrl, err := url.Parse("https://cloud.mindsdb.com")
	api.Url = apiUrl
	api.Email = "user@mail.com"
	api.Password = "**password**"
	session := &http.Client{
		Timeout: time.Second * 10,
	}
	api.Session = session
	req, err := api.Login()
	Handle_error(err)
	server := mindsdb_go_sdk.Server{}
	server = *server.NewServer(&api)
	projects := server.ListProjects()
	for _, project := range projects {
		fmt.Println(project.Name)
	}
	logged_session := req.Session
	_, _, err = api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	Handle_error(err)
}
