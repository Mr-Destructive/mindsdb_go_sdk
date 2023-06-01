package mindsdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

type Server struct {
	Api *connectors.RestAPI
}

func NewServer(api *connectors.RestAPI) *Server {
	return &Server{
		Api: api,
	}
}

const listProjectsQuery = `SELECT NAME FROM information_schema.databases;`

func Connect(apiUrl, email, password string) (*connectors.RestAPI, error) {
	api := connectors.RestAPI{}
	hostUrl := apiUrl
	if apiUrl == "" {
		hostUrl = "https://cloud.mindsdb.com"
	}
	Url, err := url.Parse(hostUrl)
	if err != nil {
		HandleError(err)
	}
	api.Url = Url
	api.Email = email
	api.Password = password

	jsonStr, _ := json.Marshal(map[string]string{
		"email":    api.Email,
		"password": api.Password,
	})
	api.Session = &http.Client{
		Timeout: time.Second * 10,
	}
	empty_api := connectors.RestAPI{}
	url := fmt.Sprintf("%s/cloud/login", api.Url.String())
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		return &empty_api, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.Session.Do(req)
	if err != nil {
		return &empty_api, err
	}
	var r connectors.Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return &empty_api, err
	}
	session_id := connectors.GetSessionId(resp)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &empty_api, fmt.Errorf("Login failed: %s", resp.Status)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: session_id,
	}
	client := &http.Client{Timeout: time.Second * 60, Jar: jar}
	client.Jar.SetCookies(api.Url, []*http.Cookie{cookie})
	api.Session = client
	return &api, nil
}

func (s *Server) ListDatabases() []connectors.ColumnType {
	_, columns, err := s.Api.SqlQuery(s.Api.Session,
		listProjectsQuery,
		"information_schema.databases",
		true,
	)
	if err != nil {
		return nil
	}
	return columns
}

func (s *Server) ListProjects() []*Project {
	query := listProjectsQuery + `WHERE TYPE = 'project';`
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	if err != nil {
		return nil
	}

	var projects []*Project
	for _, project := range data {
		for _, field := range project.Fields {
			if field != nil {
				projects = append(projects, &Project{Name: field.(string), Api: s.Api, Server: s})
			}
		}
	}
	return projects
}

func (s *Server) GetDatabase(name string) (*Database, error) {
	query := fmt.Sprintf("SHOW TABLES FROM %s;", name)
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	if err != nil {
		return nil, err
	}

	database := Database{Name: name, Api: s.Api, Server: s, Tables: data}
	return &database, nil
}

func (s *Server) CreateProject(name string) (*Project, error) {
	return NewProject(s, name, "mindsdb", map[string]string{}), nil
}

func (s *Server) DropProject(name string) error {
	query := fmt.Sprintf("DROP DATABASE %s;", name)
	_, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	return err
}

func (s *Server) GetProject(name string) (*Project, error) {
	database, err := s.GetDatabase(name)
	if err != nil {
		return nil, err
	}

	project := &Project{Name: database.Name, Server: s, Api: s.Api}
	return project, nil
}
