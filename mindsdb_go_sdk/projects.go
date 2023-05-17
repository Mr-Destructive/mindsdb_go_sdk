package mindsdb_go_sdk

import (
	"fmt"
	"mindsdb_go_sdk/connectors"
	"strings"
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

func NewProject(server *Server, name string, engine string, params map[string]string) *Project {
	parameters := ""
	for k, v := range params {
		parameters += fmt.Sprintf(`"%s": "%s",`, k, v)
	}
	parameters = strings.TrimSuffix(parameters, ",")

	query := fmt.Sprintf(`CREATE DATABASE %s WITH ENGINE = "%s", PARAMETERS = {%s};`, name, engine, parameters)

	data, columns, err := server.Api.SqlQuery(server.Api.Session, query, "", true)
	fmt.Println(data, columns)
	HandleError(err)

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
