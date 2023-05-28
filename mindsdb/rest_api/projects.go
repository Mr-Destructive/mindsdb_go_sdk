package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetProjects(api *connectors.RestAPI) ([]byte, error) {
	resp, err := api.APIRequest("", "GET", "/api/projects", map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetProject(api *connectors.RestAPI, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s", name)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
