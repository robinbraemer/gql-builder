package gqlbuilder

import (
	"github.com/graphql-go/graphql/language/ast"
	gql "github.com/graphql-go/graphql"
	"time"
	"fmt"
)

const layout = "2001-03-28T12:00:01Z"

var dateTime = gql.NewScalar(gql.ScalarConfig{
	Name:        "DateTime",
	Description: `Scala type of an ISO-8601 encoded UTC date and time string. e.g. "2001-03-28T12:00:01Z"`,
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case time.Time:
			return value.Format(layout)
		case *time.Time:
			v := *value
			return v.Format(layout)
		default:
			panic("could not serialize datetime to string")
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			t, err := time.Parse(layout, value)
			if err != nil {
				panic(fmt.Sprintf("error parsing datetime: %v", err))
			}
			return t
		case *string:
			t, err := time.Parse(layout, *value)
			if err != nil {
				panic(fmt.Sprintf("error parsing datetime: %v", err))
			}
			return t
		default:
			panic("could not parse string to datetime")
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			t, err := time.Parse(layout, valueAST.Value)
			if err != nil {
				panic("could not parse literal to datetime")
			}
			return t
		default:
			panic("could not parse literal to datetime")
		}
	},
})

func DateTimeScalar() *gql.Scalar {
	return dateTime
}
