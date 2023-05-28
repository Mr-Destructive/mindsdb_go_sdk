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
