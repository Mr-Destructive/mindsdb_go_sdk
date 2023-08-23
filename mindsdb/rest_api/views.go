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

func CreateView(api *connectors.RestAPI, project string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/views", project)
	resp, err := api.APIRequest("", "POST", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func UpdateView(api *connectors.RestAPI, project, view map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/views/%s", project, view)
	resp, err := api.APIRequest("", "PUT", endpoint, view)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteView(api *connectors.RestAPI, project, view string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/projects/%s/views/%s", project, view)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
