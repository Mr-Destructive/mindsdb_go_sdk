package rest_api

import (
	"fmt"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

const base_url = "/api/project"

func GetSkills(api *connectors.RestAPI, project string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/skills", base_url, project)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetSkill(api *connectors.RestAPI, project, skill string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/skills/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func CreateSkill(api *connectors.RestAPI, project string, body map[string]string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/skills", base_url, project)
	resp, err := api.APIRequest("", "POST", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func UpdateSkills(api *connectors.RestAPI, project, skill string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/%s/skills/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "PUT", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteSkill(api *connectors.RestAPI, project, skill string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/%s/skills/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetAgents(api *connectors.RestAPI, project string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/agents", base_url, project)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func GetAgent(api *connectors.RestAPI, project, skill string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/agents/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "GET", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}

func CreateAgent(api *connectors.RestAPI, project string, body map[string]string) ([]byte, error) {
    endpoint := fmt.Sprintf("%s/%s/agents", base_url, project)
	resp, err := api.APIRequest("", "POST", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func UpdateAgents(api *connectors.RestAPI, project, skill string, body map[string]string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/%s/agents/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "PUT", endpoint, body)
	mindsdb.HandleError(err)
	return resp, err
}

func DeleteAgent(api *connectors.RestAPI, project, skill string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/%s/agents/%s", base_url, project, skill)
	resp, err := api.APIRequest("", "DELETE", endpoint, map[string]string{})
	mindsdb.HandleError(err)
	return resp, err
}
