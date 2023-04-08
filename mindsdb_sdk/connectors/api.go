package connectors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type RestAPI struct {
	Url      *url.URL
	Email    string
	Password string
	Session  *http.Client
}

type Response struct {
	Type  string      `json:"type"`
	Error string      `json:"error_message"`
	Data  interface{} `json:"data"`
}

func getSessionId(resp *http.Response) string {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session" {
			return cookie.Value
		}
	}
	return ""
}

func (api *RestAPI) Login() (RestAPI, error) {
	jsonStr, _ := json.Marshal(map[string]string{
		"email":    api.Email,
		"password": api.Password,
	})
	empty_api := RestAPI{}
	url := fmt.Sprintf("%s/cloud/login", api.Url.String())
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		return empty_api, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.Session.Do(req)
	if err != nil {
		return empty_api, err
	}
	var r Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return empty_api, err
	}
	session_id := getSessionId(resp)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return empty_api, fmt.Errorf("Login failed: %s", resp.Status)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: session_id,
	}
	client := &http.Client{Timeout: time.Second * 10, Jar: jar}
	client.Jar.SetCookies(api.Url, []*http.Cookie{cookie})
	api.Session = client
	return *api, nil
}

func (api *RestAPI) SqlQuery(session *http.Client, sql string, database string, lowercaseColumns bool) (data [][]interface{}, columns []string, err error) {
	if database == "" {
		database = "mindsdb"
	}
	url := fmt.Sprintf("%s/api/sql/query", api.Url.String())
	jsonStr, _ := json.Marshal(map[string]interface{}{
		"query":   sql,
		"context": map[string]string{"db": database},
	})

	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := session
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println("body: ", string(body))
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("API call failed with status code %d", resp.StatusCode)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, nil, err
	}

	if dataArr, ok := result["data"].([]interface{}); ok {
		for _, row := range dataArr {
			if rowMap, ok := row.(map[string]interface{}); ok {
				rowData := make([]interface{}, len(rowMap))
				i := 0
				for _, value := range rowMap {
					rowData[i] = value
					i++
				}
				data = append(data, rowData)
			}
		}
	}

	if columnsArr, ok := result["columns"].([]interface{}); ok {
		for _, column := range columnsArr {
			if columnStr, ok := column.(string); ok {
				if lowercaseColumns {
					columns = append(columns, strings.ToLower(columnStr))
				} else {
					columns = append(columns, columnStr)
				}
			}
		}
	}

	return data, columns, nil
}
