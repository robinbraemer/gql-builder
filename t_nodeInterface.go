package gqlbuilder

import gql "github.com/graphql-go/graphql"

var nodeInterface = gql.NewInterface(gql.InterfaceConfig{
	Name:        "Node",
	Description: "An object with an ID.",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type:        gql.NewNonNull(gql.ID),
			Description: "ID of the object.",
		},
		//"test": &gql.Field{
		//	Type: gql.String,
		//},
	},
	// TODO get object by id
	// ResolveType:
})

func NodeInterface() *gql.Interface {
	return nodeInterface
}