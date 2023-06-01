package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func SqlQuery(api *connectors.RestAPI, query string) ([]byte, error) {
	body := map[string]string{"query": query}
	resp, err := api.APIRequest("", "POST", "/api/sql/query", body)
	mindsdb.HandleError(err)
    return resp, err
}

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
