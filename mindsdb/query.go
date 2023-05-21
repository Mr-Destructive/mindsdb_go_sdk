package mindsdb

import (
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

type Query struct {
	Api       *connectors.RestAPI
	Sql       string
	DBName    string
	ResultSet ResultSet
}

func MindsDBClient(baseURL string) *connectors.RestAPI {
	return &connectors.RestAPI{}
}

func NewQuery(api *connectors.RestAPI, sql string, dbName string) *Query {
	return &Query{
		Api:    api,
		Sql:    sql,
		DBName: dbName,
	}
}

//func (q *Query) Fetch() string {
//	request_body := fmt.Sprintf(`{"sql": "%s", "database": "%s"}`, q.Sql, q.DBName)
//	body := bytes.NewBuffer([]bytes(request_body))
//	resp, er := http.Post(fmt.Sprintf("%s/api/sql/query", q.api.baseURL), "application/json", body)
//	Handle_error(err)
//	defer resp.Body.Close()
//	return respBody.String()
//}
