package mindsdb_go_sdk

import (
	"mindsdb_go_sdk/connectors"
)

type View struct {
}

type Model struct {
}

type ModelVersion struct {
}

type Project struct {
	Name   string
	Server *Server
	Api    *connectors.RestAPI
}

func NewProject(server *Server, name string) *Project {
	return &Project{
		Name:   name,
		Server: server,
		Api:    server.Api,
	}
}

func (p *Project) Query(sql string) *Query {
	// Execute query and return Query object
	return &Query{}
}
