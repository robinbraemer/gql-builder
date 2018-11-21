package gqlbuilder

import (
	gql "github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var (
	EdgesDesc      = "A list of edges."
	NodesDesc      = "A list of nodes."
	PageInfoDesc   = "Information to aid in pagination."
	TotalCountDesc = "Identifies the total count of items in the connection."
)

type ConnectionConfig struct {
	// The type which will be listed/paged then.
	Type *gql.Object

	// Optional: If nil the edge type is generated.
	TypeEdge *gql.Object
	// Optional: the connection's name
	Name string
	// Optional: the connection's description
	Description string
}

func NewConnection(config ConnectionConfig) *gql.Object {
	if config.Type == nil {
		if len(config.Name) != 0 {
			log.Fatalf("error creating connection type %s: type not provided", config.Name)
		} else {
			log.Fatal("error creating a connection type: type not provided")
		}
	}

	conf := gql.ObjectConfig{}

	if len(config.Name) != 0 {
		conf.Name = config.Name
	} else {
		conf.Name = config.Type.Name() + "Connection"
	}

	if len(config.Description) != 0 {
		conf.Description = config.Description
	} else {
		conf.Description = fmt.Sprintf("The connection type for %s.", config.Type.Name())
	}

	fields := gql.Fields{}

	if config.TypeEdge != nil {
		fields["edges"] = &gql.Field{
			Type:        config.TypeEdge,
			Description: EdgesDesc,
		}
	} else {
		fields["edges"] = &gql.Field{
			Type:        NewEdge(EdgeConfig{Type: config.Type}),
			Description: EdgesDesc,
		}
	}

	fields["nodes"] = &gql.Field{
		Type:        gql.NewNonNull(gql.NewList(config.Type)),
		Description: NodesDesc,
	}
	fields["pageInfo"] = &gql.Field{
		Type:        gql.NewNonNull(PageInfoType()),
		Description: PageInfoDesc,
	}
	fields["totalCount"] = &gql.Field{
		Type:        gql.NewNonNull(gql.Int),
		Description: TotalCountDesc,
	}

	conf.Fields = fields

	return gql.NewObject(conf)
}
