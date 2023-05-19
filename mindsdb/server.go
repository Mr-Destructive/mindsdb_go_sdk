package mindsdb

import (
	"fmt"
	"mindsdb/connectors"
)

type Server struct {
	Api *connectors.RestAPI
}

func (s *Server) NewServer(api *connectors.RestAPI) *Server {
	server := &Server{
		Api: api,
	}
	return server
}

func (s *Server) listDatabases() []string {
	_, columns, err := s.Api.SqlQuery(s.Api.Session,
		"select NAME from information_schema.databases",
		"information_schema.databases",
		true,
	)
	HandleError(err)
	return columns
}

func (s *Server) listProjects() []string {
	data, columns, err := s.Api.SqlQuery(s.Api.Session,
		"select NAME from information_schema.databases where TYPE='project'",
		"",
		true,
	)
	HandleError(err)
	var projects []string
	for _, project := range data {
		for column_idx, _ := range columns {
			projects = append(projects, project.Fields[column_idx].(string))
		}
	}
	return projects
}

func (s *Server) GetDatabase(name string) (*Database, error) {
	query := fmt.Sprintf(`SHOW TABLES FROM %s;`, name)
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	database := Database{Name: name, Api: s.Api, Server: s, Tables: data}
	HandleError(err)
	return &database, nil
}

func (s *Server) ListProjects() []*Project {
	names := s.listProjects()
	projects := make([]*Project, len(names))
	for i, name := range names {
		projects[i] = &Project{Name: name, Api: s.Api, Server: s}
	}
	return projects
}

func (s *Server) CreateProject(name string) (*Project, error) {
	return NewProject(s, name, "mindsdb", map[string]string{}), nil
}

func (s *Server) DropProject(name string) string {
	query := fmt.Sprintf(`DROP DATABASE %s;`, name)
	_, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	return "Database deleted successfully"
}

func (s *Server) GetProject(name string) (*Project, error) {
	database, err := s.GetDatabase(name)
	project := &Project{Name: database.Name, Server: s, Api: s.Api}
	HandleError(err)
	return project, nil
}
