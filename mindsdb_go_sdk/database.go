package mindsdb_go_sdk

import (
	"fmt"
	"mindsdb_go_sdk/connectors"
	"strings"
)

type Database struct {
	Server *Server
	Name   string
	Api    *connectors.RestAPI
}

func NewDatabase(server *Server, name string) *Database {
	return &Database{
		Server: server,
		Name:   name,
		Api:    server.Api,
	}
}

func (d *Database) String() string {
	return ""
}

func (d *Database) DBQuery(sql string) *Query {
	return NewQuery(d.Api, sql, d.Name)
}

func (d *Database) listTables() []string {
	// Retrieve list of table names
	return []string{}
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
		// fixme: schemas not visible in 'show tables'
		return nil, fmt.Errorf("table doesn't exist")
	}
	return NewTable(d, name), nil
}

func (d *Database) CreateTable(name string, query *Query, replace bool) (*Table, error) {
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}
	// Create or replace table based on the query
	return NewTable(d, name), nil
}

type Table struct {
	// Table fields
}

func NewTable(database *Database, name string) *Table {
	return &Table{
		// Initialize Table fields
	}
}
