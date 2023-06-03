package mindsdb

import (
	"fmt"
	"strings"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

type Model struct {
	Name            string
	Engine          string
	Project         *Project
	Version         float64
	Status          string
	Accuracy        float64
	PredictColumn   string
	UpdateStatus    string
	MindsdbVersion  string
	TrainingOptions map[string]interface{}
	Parameters      map[string]string
}

type ResultSet struct {
	Columns []connectors.ColumnType
	Rows    []connectors.Record
}

func convertInterfaceToString(value interface{}) (string, error) {
	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("unable to convert interface to string")
}

func convertInterfaceToFloat64(value interface{}) (float64, error) {
	if num, ok := value.(float64); ok {
		return num, nil
	}
	return 0.0, fmt.Errorf("unable to convert interface to float64")
}

func (s *Server) ListModels(project string) ([]Model, error) {
	query := fmt.Sprintf(`SHOW MODELS FROM %s;`, project)
	data, _, err := s.Api.SqlQuery(s.Api.Session, query, "", true)
	if err != nil {
		return nil, err
	}

	var models []Model
	for _, d := range data {
		result := d
		for i, row := range d.Fields {
			if row == nil && i == 5 {
				result.Fields[i] = 0.0
			}
		}

		model := Model{
			Name:            getStringValue(result.Fields[0]),
			Engine:          getStringValue(result.Fields[1]),
			Project:         &Project{Name: getStringValue(result.Fields[2]), Api: s.Api, Server: s},
			Version:         getFloat64Value(result.Fields[3]),
			Status:          getStringValue(result.Fields[4]),
			Accuracy:        getFloat64Value(result.Fields[5]),
			PredictColumn:   getStringValue(result.Fields[6]),
			TrainingOptions: getMapValue(result.Fields[11]),
		}

		models = append(models, model)
	}

	return models, nil
}

func getStringValue(value interface{}) string {
	str, _ := convertInterfaceToString(value)
	return str
}

func getFloat64Value(value interface{}) float64 {
	num, _ := convertInterfaceToFloat64(value)
	return num
}

func getMapValue(value interface{}) map[string]interface{} {
	if m, ok := value.(map[string]interface{}); ok {
		return m
	}
	return nil
}

func (m *Model) GetStatus() bool {
	query := fmt.Sprintf(`SELECT status FROM %s.models WHERE name = '%s';`, m.Project.Name, m.Name)
	result, err := m.Project.Query(query, m.Project.Name)
	HandleError(err)
	if result.ResultSet.Rows[0].Fields[0].(string) == "complete" {
		return true
	}
	return false

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

func (m *Model) Retrain() (*Model, error) {
	if m.UpdateStatus == "available" {
		query := fmt.Sprintf(`RETRAIN %s.%s PREDICT %s USING ENGINE = '%s', %s;`, m.Project.Name, m.Name, m.PredictColumn, m.Engine, m.Parameters)
		_, err := m.Project.Query(query, m.Project.Name)
		if err != nil {
			return nil, err
		}
		model, err := m.Project.GetModel(m.Name)
		if err != nil {
			return nil, err
		}
		return model, nil
	}
	return nil, nil
}

func (m *Model) FineTune(db, query string) (*Model, error) {
	finetune_query := fmt.Sprintf(`FINETUNE %s.%s FROM %s (%s);`, m.Project.Name, m.Name, db, query)
	_, err := m.Project.Query(finetune_query, m.Project.Name)
	if err != nil {
		return nil, err
	}
	model, err := m.Project.GetModel(m.Name)
	if err != nil {
		return nil, err
	}
	return model, nil
}
