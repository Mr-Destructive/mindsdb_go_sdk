package mindsdb_go_sdk

import (
	"fmt"
	"mindsdb_go_sdk/connectors"
	"strings"
)

type Model struct {
	Name            string
	Engine          string
	Project         *Project
	Version         string
	Status          bool
	Accuracy        float64
	Predict         string
	UpdateStatus    string
	MindsdbVersion  string
	TrainingOptions map[string]string
	SelectDataQuery string
}

type ResultSet struct {
	Columns []string
	Rows    []connectors.Record
}

func (s *Server) NewModel(project string, modelName string, predictColumn string, engine string, params map[string]string) *Model {
	parameters := ""
	for k, v := range params {
		parameters += fmt.Sprintf(`%s = '%s',`, k, v)
	}
	parameters = strings.TrimSuffix(parameters, ",")

	query := fmt.Sprintf(`CREATE MODEL %s.%s PREDICT %s USING ENGINE = '%s', %s;`, project, modelName, predictColumn, engine, parameters)
	data, columns, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	fmt.Println(data, columns)
	model := &Model{
		Name: params["name"],
	}

	return model
}

func (s *Server) GetModel(project string, name string) *Model {
	query := fmt.Sprintf(`SELECT * FROM %s.models WHERE name='%s';`, project, name)
	data, columns, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	fmt.Println(data, columns)
	return &Model{}
}

func (s *Server) Predict(project string, name string, predictColumn string, params map[string]string) *ResultSet {
	parameters := ""
	for k, v := range params {
		parameters += fmt.Sprintf(`%s = '%s' AND `, k, v)
	}
	parameters = strings.TrimSuffix(parameters, "AND ")
	fmt.Println(parameters)
	query := fmt.Sprintf(`SELECT %s FROM %s.%s WHERE %s;`, predictColumn, project, name, parameters)
	fmt.Println(query)
	data, columns, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	resultSet := ResultSet{Columns: columns, Rows: data}
	return &resultSet
}
