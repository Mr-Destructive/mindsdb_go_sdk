package mindsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
	"strings"
)

type View struct {
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

func (p *Project) Query(sql, database string) *Query {
	if database == "" {
		database = "mndsdb"
	}
	data, column, err := p.Api.SqlQuery(p.Api.Session, sql, database, true)
	HandleError(err)
	resultSet := ResultSet{Columns: column, Rows: data}
	return &Query{Api: p.Api, Sql: sql, DBName: database, ResultSet: resultSet}
}

func (p *Project) NewModel(modelName string, predictColumn string, engine string, params map[string]string) (*Model, error) {
	parameters := ""
	for k, v := range params {
		parameters += fmt.Sprintf(`%s = '%s',`, k, v)
	}
	parameters = strings.TrimSuffix(parameters, ",")

	query := fmt.Sprintf(`CREATE MODEL %s.%s PREDICT %s USING ENGINE = '%s', %s;`, p.Name, modelName, predictColumn, engine, parameters)
	data, _, err := p.Api.SqlQuery(p.Api.Session, query, "", true)
	HandleError(err)
	if len(data) > 0 {
		result := data[0]
		model := Model{
			Name:    result.Fields[0].(string),
			Engine:  result.Fields[1].(string),
			Project: p,
		}

		var conversionErr error
		if model.Version, conversionErr = convertInterfaceToFloat64(result.Fields[3]); conversionErr != nil {
			return nil, conversionErr
		}

		if model.Status, conversionErr = convertInterfaceToString(result.Fields[4]); conversionErr != nil {
			return nil, conversionErr
		}

		if model.Accuracy, conversionErr = convertInterfaceToFloat64(result.Fields[5]); conversionErr != nil {
			return nil, conversionErr
		}
		if model.PredictColumn, conversionErr = convertInterfaceToString(result.Fields[6]); conversionErr != nil {
			return nil, conversionErr
		}

		if trainingOptions, ok := result.Fields[12].(map[string]interface{}); ok {
			model.TrainingOptions = trainingOptions
		} else if jsonString, ok := result.Fields[12].(string); ok {
			options := make(map[string]interface{})
			jsonString = strings.ReplaceAll(jsonString, "'", `"`)
			err := json.Unmarshal([]byte(jsonString), &options)
			if err != nil {
				panic(err)
			}
			model.TrainingOptions = options
		}

		return &model, nil
	}
	return nil, errors.New("Model already exists or failed to create Model")
}

func (p *Project) GetModel(name string) *Model {
	query := fmt.Sprintf(`SELECT * FROM %s.models WHERE name='%s';`, p.Name, name)
	data, _, err := p.Api.SqlQuery(p.Api.Session, query, "", true)
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
		Project:         p,
		Version:         result.Fields[3].(float64),
		Status:          result.Fields[4].(string),
		Accuracy:        result.Fields[5].(float64),
		PredictColumn:   result.Fields[6].(string),
		TrainingOptions: result.Fields[11].(map[string]interface{}),
	}
	return &model
}

func (m *Model) Predict(predictColumn string, params map[string]string) *ResultSet {
	parameters := ""
	for k, v := range params {
		parameters += fmt.Sprintf(`%s = '%s' AND `, k, v)
	}
	parameters = strings.TrimSuffix(parameters, "AND ")
	query := fmt.Sprintf(`SELECT %s FROM %s.%s WHERE %s;`, predictColumn, m.Project.Name, m.Name, parameters)
	data, columns, err := m.Project.Api.SqlQuery(m.Project.Api.Session, query, "", true)
	HandleError(err)
	resultSet := ResultSet{Columns: columns, Rows: data}
	return &resultSet
}

func (p *Project) DropModel(name string) string {
	query := fmt.Sprintf(`DROP MODEL %s.%s;`, p.Name, name)
	_, _, err := p.Api.SqlQuery(p.Api.Session, query, "", true)
	HandleError(err)
	return "Model deleted successfully"
}
