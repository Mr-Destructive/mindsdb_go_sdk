package mindsdb

import (
	"fmt"
	"mindsdb/connectors"
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
}

type ResultSet struct {
	Columns []string
	Rows    []connectors.Record
}

func convertInterfaceToString(value interface{}) (string, error) {
	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", nil
}

func convertInterfaceToFloat64(value interface{}) (float64, error) {
	if num, ok := value.(float64); ok {
		return num, nil
	}
	return 0.0, nil
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
			PredictColumn:   result.Fields[6].(string),
			TrainingOptions: result.Fields[11].(map[string]interface{}),
		}
		models = append(models, model)
	}
	return models
}
