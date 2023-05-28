package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetDatabases(api *connectors.RestAPI) ([]byte, error) {
	resp, err := api.APIRequest("", "GET", "/api/databases", map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetDatabase(api *connectors.RestAPI, database string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s", database)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
