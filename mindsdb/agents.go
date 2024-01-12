package mindsdb

import (

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
)

type Skill struct {
	Name        string
	Type        string
	Source      string
	Description string
	Server      *Server
	Api         *connectors.RestAPI
}

func NewSkill(server *Server, name, stype, source, description string) *Skill {
	return &Skill{
        Name:        name,
        Type:        stype,
        Source:      source,
        Description: description,
        Server:      server,
        Api:         server.Api,
	}
}

