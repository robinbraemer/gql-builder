package main

import (
	gql "github.com/graphql-go/graphql"
	gqlh "github.com/graphql-go/handler"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Animal struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var animals []Animal
var animals2 []Animal

func init() {
	animals = append(animals, Animal{Name: "Cow", Id: "1"}, Animal{Id: "3", Name: "Dog"})
	animals2 = append(animals, Animal{Name: "Chicken", Id: "2"}, Animal{Name: "Chicken", Id: "4"}, Animal{Name: "Chicken", Id: "5"})
}

var animalType = gql.NewObject(gql.ObjectConfig{
	Name: "Animal",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.ID,
		},
		"name": &gql.Field{
			Type: gql.String,
		},
	},
})

var animalSchema, _ = gql.NewSchema(gql.SchemaConfig{
	Query: gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"animals": &gql.Field{
				Type: gql.NewNonNull(gql.NewObject(gql.ObjectConfig{
					Name: "AnimalConnection",
					Fields: gql.Fields{
						"nodes": &gql.Field{
							Type: gql.NewNonNull(gql.NewList(animalType)),
							Resolve: func(p gql.ResolveParams) (interface{}, error) {
								fmt.Println("resolver nodes")
								return animals2, nil
							},
						},
					},
				})),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					fmt.Println("resolver animals")
					return animals, nil
				},
			},
		},
	}),
})

func main() {

	// HTTP GraphQL handler
	h := gqlh.New(&gqlh.Config{
		Schema:   &animalSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	r := gin.New()
	r.POST("", gin.WrapH(h))
	r.GET("", gin.WrapH(h))

	fmt.Printf("Graphql server listening on %s.\n", "http://127.0.0.1:8080")
	fmt.Println(r.Run("127.0.0.1:8080"))
}
