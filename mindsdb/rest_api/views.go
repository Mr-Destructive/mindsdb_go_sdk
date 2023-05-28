package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetViews(api *connectors.RestAPI, project string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/views", project)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetView(api *connectors.RestAPI, project, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/views/%s", project, name)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
