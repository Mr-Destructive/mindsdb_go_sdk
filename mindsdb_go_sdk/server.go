package mindsdb_go_sdk

import (
	"fmt"
	"mindsdb_go_sdk/connectors"
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
			projects = append(projects, project.Fields[column_idx])
		}
	}
	return projects
}

func (s *Server) GetDatabase(name string) (*Database, error) {
	names := s.listDatabases()
	for _, n := range names {
		if n == name {
			return NewDatabase(s, name), nil
		}
	}
	return nil, fmt.Errorf("database doesn't exist")
}

func (s *Server) ListProjects() []*Project {
	names := s.listProjects()
	projects := make([]*Project, len(names))
	for i, name := range names {
		projects[i] = NewProject(s, name)
	}
	return projects
}

func (s *Server) CreateProject(name string) (*Project, error) {
	return NewProject(s, name), nil
}

func (s *Server) DropProject(name string) {
}

func (s *Server) GetProject(name string) (*Project, error) {
	names := s.ListProjects()
	for _, n := range names {
		if n.Name == name {
			return NewProject(s, name), nil
		}
	}
	return nil, fmt.Errorf("project doesn't exist")
}
