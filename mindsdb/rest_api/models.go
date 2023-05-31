package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetModels(api *connectors.RestAPI, project string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models", project)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetModel(api *connectors.RestAPI, project, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models/%s", project, name)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func DescribeModel(api *connectors.RestAPI, project, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models/%s/describe", project, name)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func QueryModel(api *connectors.RestAPI, project, model string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models/%s", project, model)
	resp, err := api.APIRequest("", "POST", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func TrainModel(api *connectors.RestAPI, project string, query map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models", project)
	resp, err := api.APIRequest("", "POST", endpoint, query)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteModel(api *connectors.RestAPI, project, model string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/models/%s", project, model)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
