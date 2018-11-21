package gqlbuilder

import (
	gql "github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var (
	CursorDesc = "A cursor for use in pagination."
	NodeDesc   = "The item at the end of the edge."
)

type EdgeConfig struct {
	// The type which will be listed/paged/cursored then.
	Type *gql.Object

	// Optional: the edge's name
	Name string
	// Optional: the edge's description
	Description string
}

func NewEdge(config EdgeConfig) *gql.Object {
	if config.Type == nil {
		if len(config.Name) != 0 {
			log.Fatalf("error creating edge type %s: type not provided", config.Name)
		} else {
			log.Fatal("error creating a edge type: type not provided")
		}
	}

	fields := gql.Fields{}
	conf := gql.ObjectConfig{
		Fields: fields,
	}

	if len(config.Name) != 0 {
		conf.Name = config.Name
	} else {
		conf.Name = config.Type.Name() + "Edge"
	}

	if len(config.Description) != 0 {
		conf.Description = config.Description
	} else {
		conf.Description = fmt.Sprintf("An edge in a connection for type %s.", config.Type.Name())
	}

	fields["cursor"] = &gql.Field{
		Type:        gql.NewNonNull(gql.String),
		Description: CursorDesc,
	}
	fields["node"] = &gql.Field{
		Type:        config.Type,
		Description: NodeDesc,
	}

	return gql.NewObject(conf)
}
