package user

import (
	gqlb "github.com/cloudfound/gql-builder"
	gql "github.com/graphql-go/graphql"
	"time"
)

var Users = []User{
	{Id: "robin", CreatedAt: time.Now(), Name: "Robin"},
	{Id: "bob", CreatedAt: time.Now(), Name: "Bob"},
}

func InitTopic() *gqlb.SchemaTopic {
	return &gqlb.SchemaTopic{
		QueryFields: &gql.Fields{
			"users": &gql.Field{
				Type: gql.NewNonNull(UserConnection()),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return Users, nil
				},
			},
		},
	}
}
