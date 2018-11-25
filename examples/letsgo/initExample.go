package letsgo

import (
	gql "github.com/graphql-go/graphql"
	gqlb "github.com/cloudfound/gql-builder"
	"time"
)

func InitTopic() *gqlb.SchemaTopic {
	queryFields := &gql.Fields{
			"lets": &gql.Field{
				Type: gql.String,
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return "Go!", nil
				},
				Description: "Hey, you!",
			},
			"go": &gql.Field{
				Type: gqlb.DateTimeScalar(),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return time.Now(), nil
				},
				Description: "Get on the code!",
			},
		}

	return &gqlb.SchemaTopic{
		QueryFields: queryFields,
	}
}
