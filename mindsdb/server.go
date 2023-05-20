package mindsdb

import (
	"fmt"
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

func (s *Server) ListDatabases() []string {
	_, columns, err := s.Api.SqlQuery(s.Api.Session,
		"SELECT NAME FROM information_schema.databases",
		"information_schema.databases",
		true,
	)
	if err != nil {
		// Handle the error appropriately
		return nil
	}
	return columns
}

func (s *Server) ListProjects() []*Project {
	query := `
		SELECT NAME FROM information_schema.databases
		WHERE TYPE = 'project';
	`
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	if err != nil {
		return nil
	}

	var projects []*Project
	for _, project := range data {
		for _, field := range project.Fields {
			projects = append(projects, &Project{Name: field.(string), Api: s.Api, Server: s})
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
