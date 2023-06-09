package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

func GetTables(api *connectors.RestAPI, database string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s/tables", database)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetTable(api *connectors.RestAPI, database, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s/tables/%s", database, name)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func CreateTable(api *connectors.RestAPI, database string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s/tables", database)
	resp, err := api.APIRequest("", "POST", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteTable(api *connectors.RestAPI, database, name string) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/databases/%s/tables/%s", database, name)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
