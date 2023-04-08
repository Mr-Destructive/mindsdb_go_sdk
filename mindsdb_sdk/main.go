package main

import (
	"log"
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
	api.Email = "your@email.com"
	api.Password = "your password"
	session := &http.Client{
		Timeout: time.Second * 10,
	}
	api.Session = session
	req, err := api.Login()
	Handle_error(err)
	logged_session := req.Session
	data, column, err := api.SqlQuery(logged_session, "SELECT NAME FROM models;", "", true)
	Handle_error(err)
	log.Println(data)
	log.Println(column)
}
