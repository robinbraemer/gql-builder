package main

import (
	gqlb "github.com/cloudfound/gql-builder"
	gqlHandler "github.com/graphql-go/handler"
	"github.com/cloudfound/gql-builder/examples/letsgo"
	"github.com/cloudfound/gql-builder/examples/user"
	"fmt"
	"syscall"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const Endpoint = "/example"

func main() {
	log.SetLevel(log.DebugLevel)
	// Build schema for internal graph API
	schema, err := gqlb.Build(gqlb.Topics(
		//
		// Extent the schema by simply adding more initializer functions.
		//
		letsgo.InitTopic,
		user.InitTopic,
	))
	if err != nil {
		fmt.Println(err.Error())
		syscall.Exit(1)
	}

	fmt.Println("Schema initialized successfully.")

	// HTTP GraphQL handler
	h := gqlHandler.New(&gqlHandler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r := gin.New()
	r.POST(Endpoint, gin.WrapH(h))
	r.GET(Endpoint, gin.WrapH(h))

	fmt.Printf("Graphql server listening on %s.\n", "http://127.0.0.1:8080"+Endpoint)
	fmt.Println(r.Run("127.0.0.1:8080"))
}
