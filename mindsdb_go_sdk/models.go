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
	Version         float64
	Status          string
	Accuracy        float64
	Predict         string
	UpdateStatus    string
	MindsdbVersion  string
	TrainingOptions map[string]interface{}
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
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	result := data[0]
	for i, row := range result.Fields {
		if row == nil && i == 5 {
			result.Fields[i] = 0.0
		}
	}
	model := Model{
		Name:            result.Fields[0].(string),
		Engine:          result.Fields[1].(string),
		Project:         &Project{Name: result.Fields[2].(string), Api: s.Api, Server: s},
		Version:         result.Fields[3].(float64),
		Status:          result.Fields[4].(string),
		Accuracy:        result.Fields[5].(float64),
		Predict:         result.Fields[6].(string),
		TrainingOptions: result.Fields[11].(map[string]interface{}),
	}
	return &model
}

func (s *Server) ListModels(project string) []Model {
	query := fmt.Sprintf(`SHOW MODELS FROM %s;`, project)
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	var models []Model
	for _, d := range data {
		result := d
		for i, row := range d.Fields {
			if row == nil && i == 5 {
				result.Fields[i] = 0.0
			}
		}
		model := Model{
			Name:            result.Fields[0].(string),
			Engine:          result.Fields[1].(string),
			Project:         &Project{Name: result.Fields[2].(string), Api: s.Api, Server: s},
			Version:         result.Fields[3].(float64),
			Status:          result.Fields[4].(string),
			Accuracy:        result.Fields[5].(float64),
			Predict:         result.Fields[6].(string),
			TrainingOptions: result.Fields[11].(map[string]interface{}),
		}
		models = append(models, model)
	}
	return models
}

func (s *Server) GetModel(project string, name string) *Model {
	query := fmt.Sprintf(`SELECT * FROM %s.models WHERE name='%s';`, project, name)
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	result := data[0]
	for i, row := range result.Fields {
		if row == nil && i == 5 {
			result.Fields[i] = 0.0
		}
	}
	model := Model{
		Name:            result.Fields[0].(string),
		Engine:          result.Fields[1].(string),
		Project:         &Project{Name: name, Api: s.Api, Server: s},
		Version:         result.Fields[3].(float64),
		Status:          result.Fields[4].(string),
		Accuracy:        result.Fields[5].(float64),
		Predict:         result.Fields[6].(string),
		TrainingOptions: result.Fields[11].(map[string]interface{}),
	}
	return &model
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

func (s *Server) DropModel(project string, name string) string {
	query := fmt.Sprintf(`DROP MODEL %s.%s;`, project, name)
	_, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	HandleError(err)
	return "Model deleted successfully"
}
