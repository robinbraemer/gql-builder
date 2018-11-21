package user

import (
	gqlb "github.com/cloudfound/gql-builder"
	gql "github.com/graphql-go/graphql"
	"time"
)

func InitTopic() *gqlb.SchemaTopic {
	return &gqlb.SchemaTopic{
		QueryFields: &gql.Fields{
			"users": usersQuery(),
		},
	}
}

var Users = []User{
	{id: "robin", createdAt: time.Now(), name: "Robin"},
	{id: "bob", createdAt: time.Now(), name: "Bob"},
}

func usersQuery() *gql.Field {
	return &gql.Field{
		Type:        gql.NewNonNull(UserConnection()),
		Description: "A list of users.",
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return Users, nil
		},
	}
}
