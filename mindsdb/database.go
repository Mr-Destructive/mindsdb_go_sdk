package mindsdb

import (
	"fmt"
	"strings"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

type Database struct {
	Name   string
	Tables []connectors.Record
	Server *Server
	Api    *connectors.RestAPI
}

func NewDatabase(server *Server, name string) *Database {
	return &Database{
		Server: server,
		Name:   name,
		Api:    server.Api,
		Tables: []connectors.Record{},
	}
}

func (d *Database) String() string {
	return ""
}

func (d *Database) Query(sql string) (*Query, error) {
	body, err := connectors.PostQuery(*d.Api, d.Name, sql)
	if err != nil {
		return &Query{}, nil
	}
	rows, columns, err := connectors.ReadRowColumns(body)
	if err != nil {
		return &Query{}, nil
	}
	return &Query{
		Api:       d.Api,
		Sql:       sql,
		DBName:    d.Name,
		ResultSet: ResultSet{Rows: rows, Columns: columns},
	}, nil
}

func (d *Database) listTables() []string {
	query := fmt.Sprintf("SHOW TABLES FROM %s;", d.Name)
	data, _, err := d.Api.SqlQuery(d.Api.Session, query, "", true)
	var tables []string
	HandleError(err)
	for _, table := range data {
		for _, field := range table.Fields {
			if field != nil {
				tables = append(tables, field.(string))
			}
		}
	}

	return tables
}

func (d *Database) ListTables() []*Table {
	names := d.listTables()
	tables := make([]*Table, len(names))
	for i, name := range names {
		tables[i] = NewTable(d, name)
	}
	return tables
}

func (d *Database) GetTable(name string) (*Table, error) {
	names := d.listTables()
	for _, n := range names {
		if n == name {
			return NewTable(d, name), nil
		}
	}
	if !strings.Contains(name, ".") {
		return nil, fmt.Errorf("table doesn't exist")
	}
	return NewTable(d, name), nil
}

func (d *Database) CreateTable(name string, query *Query, replace bool) (*Table, error) {
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}
	return NewTable(d, name), nil
}

type Table struct {
	Name string
}

func NewTable(database *Database, name string) *Table {
	return &Table{Name: name}
}
