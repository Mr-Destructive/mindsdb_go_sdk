package rest_api

import (
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetProjects(api *connectors.RestAPI) ([]byte, error) {
	resp, err := api.APIRequest("", "GET", "/api/projects", map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
