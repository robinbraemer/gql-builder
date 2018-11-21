package gqlbuilder

import gql "github.com/graphql-go/graphql"

var pageInfo = gql.NewObject(gql.ObjectConfig{
	Name:        "PageInfo",
	Description: "Information about pagination in a connection.",
	Fields: gql.Fields{
		"endCursor": &gql.Field{
			Type:        gql.String,
			Description: "When paginating forwards, the cursor to continue.",
		},
		"startCursor": &gql.Field{
			Type:        gql.String,
			Description: "When paginating backwards, the cursor to continue.",
		},
		"hasNextPage": &gql.Field{
			Type:        gql.NewNonNull(gql.Boolean),
			Description: "When paginating forwards, are there more items?",
		},
		"hasPreviousPage": &gql.Field{
			Type:        gql.NewNonNull(gql.Boolean),
			Description: "When paginating backwards, are there more items?",
		},
	},
})

func PageInfoType() *gql.Object {
	return pageInfo
}