package user

import (
	"time"
	gql "github.com/graphql-go/graphql"
	gqlb "github.com/cloudfound/gql-builder"
)

type User struct {
	// The global unique permanent id of the user.
	id string

	// The name of the user.
	name string

	// When the account was created.
	createdAt time.Time
}

var userType = gql.NewObject(gql.ObjectConfig{
	Name:        "User",
	Description: "A user is an individual's account.",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.ID),
		},

		"createdAt": &gql.Field{
			Type:        gqlb.DateTimeScalar(),
			Description: "When the account was created.",
		},

		"name": &gql.Field{
			Type:        gql.String,
			Description: "Name of the user.",
		},
	},
	Interfaces: []*gql.Interface{
		gqlb.NodeInterface(),
	},
})

func UserType() *gql.Object {
	return userType
}

var userConnection = gqlb.NewConnection(gqlb.ConnectionConfig{
	Type: UserType(),
})

func UserConnection() *gql.Object {
	return userConnection
}
