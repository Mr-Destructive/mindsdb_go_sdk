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

type Record struct {
	Fields []interface{}
}

type ColumnType struct {
	ColumnName string
	ColumnType string
}

type Response struct {
	Type  string      `json:"type"`
	Error string      `json:"error_message"`
	Data  interface{} `json:"data"`
}

func GetSessionId(resp *http.Response) string {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session" {
			return cookie.Value
		}
	}
	return ""
}

func Login(email, password string) (*RestAPI, error) {
	api := RestAPI{}
	apiUrl, err := url.Parse("https://cloud.mindsdb.com")
	api.Url = apiUrl
	api.Email = email
	api.Password = password

	jsonStr, _ := json.Marshal(map[string]string{
		"email":    api.Email,
		"password": api.Password,
	})
	api.Session = &http.Client{
		Timeout: time.Second * 10,
	}
	empty_api := RestAPI{}
	url := fmt.Sprintf("%s/cloud/login", api.Url.String())
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		return &empty_api, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.Session.Do(req)
	if err != nil {
		return &empty_api, err
	}
	var r Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return &empty_api, err
	}
	session_id := GetSessionId(resp)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &empty_api, fmt.Errorf("Login failed: %s", resp.Status)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: session_id,
	}
	client := &http.Client{Timeout: time.Second * 60, Jar: jar}
	client.Jar.SetCookies(api.Url, []*http.Cookie{cookie})
	api.Session = client
	return &api, nil
}

func (api *RestAPI) SqlQuery(session *http.Client, sql string, database string, lowercaseColumns bool) (data []Record, columns []ColumnType, err error) {
	if database == "" {
		database = "mindsdb"
	}
	body, err := PostQuery(*api, database, sql)
	return ReadRowColumns(body)
}

func PostQuery(api RestAPI, database, sql string) ([]byte, error) {

	url := fmt.Sprintf("%s/api/sql/query", api.Url.String())
	jsonStr, err := json.Marshal(map[string]interface{}{
		"query":   sql,
		"context": map[string]string{"db": database},
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := api.Session
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status code %d", resp.StatusCode)
	}
	return body, nil
}

func ReadRowColumns(body []byte) (data []Record, columns []ColumnType, err error) {

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, nil, err
	}
	var columnNames []string
	if columnsArr, ok := result["column_names"].([]interface{}); ok {
		for _, column := range columnsArr {
			if columnStr, ok := column.(string); ok {
				columnNames = append(columnNames, strings.ToLower(columnStr))
			}
		}
	}

	var records []Record
	if dataArr, ok := result["data"].([]interface{}); ok {
		for i, row := range dataArr {
			if rowMap, ok := row.([]interface{}); ok {
				record := Record{
					Fields: make([]interface{}, len(columns)),
				}
				for j, value := range rowMap {
					record.Fields = append(record.Fields, value)
					if i == 0 {
						col := ColumnType{ColumnName: columnNames[j], ColumnType: fmt.Sprintf("%T", record.Fields[j])}
						columns = append(columns, col)
					}
				}
				records = append(records, record)
			}
		}
	}
	return records, columns, nil
}
