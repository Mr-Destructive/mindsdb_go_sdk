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

func CreateDatabase(api *connectors.RestAPI, body map[string]string) ([]byte, error) {
	resp, err := api.APIRequest("", "POST", "/api/databases", body)
	mindsdb.HandleError(err)
	return resp, err
}

func UpdateDatabase(api *connectors.RestAPI, database string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s", database)
	resp, err := api.APIRequest("", "PUT", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteDatabase(api *connectors.RestAPI, database string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s", database)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
