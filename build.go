package gqlbuilder

import (
	gql "github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	RootQuery        = "Query"
	RootMutation     = "Mutation"
	RootSubscription = "Subscription"

	RootQueryDescription        = "The query root."
	RootMutationDescription     = "The mutation root."
	RootSubscriptionDescription = "The subscription root."
)

// Shortcut for *[]func() *
func Topics(topics ...func() *SchemaTopic) *[]func() *SchemaTopic {
	return &topics
}

// This is the entry for building a whole schema.
func Build(topics *[]func() *SchemaTopic) (*gql.Schema, error) {
	return buildSchema(topics)
}

type InitTopic interface {
	InitTopic() func() *SchemaTopic
}

type SchemaTopic struct {
	QueryFields        *gql.Fields
	MutationFields     *gql.Fields
	SubscriptionFields *gql.Fields

	//QueryInterfaces        []*gql.Interface
	//MutationInterfaces     []*gql.Interface
	//SubscriptionInterfaces []*gql.Interface

	Types []*gql.Type

	Directives []*gql.Directive
}

// Internal build func
func buildSchema(schemaTopicInitializer *[]func() *SchemaTopic) (*gql.Schema, error) {
	topics := *schemaTopicInitializer

	// Collect all SchemaTopics
	allQueryFields := make(gql.Fields, 5)
	allMutationFields := make(gql.Fields, 5)
	allSubscriptionFields := make(gql.Fields, 5)

	allTypes := make([]gql.Type, 0, 5)
	allDirectives := make([]*gql.Directive, 0, 5)

	//

	log.Debugf("%d schema topics found.", len(topics))
	for _, getSchemaTopic := range topics {
		schemaTopic := getSchemaTopic()
		if schemaTopic == nil {
			continue
		}

		// fields
		if schemaTopic.QueryFields != nil {
			for name, field := range *schemaTopic.QueryFields {
				allQueryFields[name] = field
			}
		}
		if schemaTopic.MutationFields != nil {
			for name, field := range *schemaTopic.MutationFields {
				allMutationFields[name] = field
			}
		}
		if schemaTopic.SubscriptionFields != nil {
			for name, field := range *schemaTopic.SubscriptionFields {
				allSubscriptionFields[name] = field
			}
		}
		// directives
		if schemaTopic.Directives != nil {
			for _, directive := range schemaTopic.Directives {
				allDirectives = append(allDirectives, directive)
			}
		}
		// types
		if schemaTopic.Types != nil {
			for _, typE := range schemaTopic.Types {
				allTypes = append(allTypes, *typE)
			}
		}
	}

	log.Debugf("%d query fields found.", len(allQueryFields))
	log.Debugf("%d mutation fields found.", len(allMutationFields))
	log.Debugf("%d subscription fields found.", len(allSubscriptionFields))

	//
	//
	// Create schema

	newRootObject := func(name, description string, fields *gql.Fields) *gql.Object {
		if len(*fields) != 0 {
			oc := &gql.ObjectConfig{}
			oc.Fields = *fields
			if len(name) != 0 {
				oc.Name = name
			}
			if len(description) != 0 {
				oc.Description = description
			}
			return gql.NewObject(*oc)
		}
		return nil
	}

	newSchema := func(query, mutation, subscription *gql.Object, types *[]gql.Type, directives *[]*gql.Directive) (gql.Schema, error) {
		schemaConf := gql.SchemaConfig{}

		if query != nil {
			schemaConf.Query = query
			log.Debug("Query root type named '" + query.Name() + "' used.")
		}
		if mutation != nil {
			schemaConf.Mutation = mutation
			log.Debug("Mutation root type named '" + mutation.Name() + "' used.")
		}
		if subscription != nil {
			schemaConf.Subscription = subscription
			log.Debug("Subscription root type named '" + subscription.Name() + "' used.")
		}

		if len(*types) != 0 {
			schemaConf.Types = *types
			log.Debug("%d types added.", len(*types))
		}
		if len(*directives) != 0 {
			schemaConf.Directives = *directives
			log.Debugf("%d directives added. %+v", len(*directives), *directives)
		}

		return gql.NewSchema(schemaConf)
	}

	schema, err := newSchema(
		newRootObject(
			RootQuery,
			RootQueryDescription,
			&allQueryFields,
		),
		newRootObject(
			RootMutation,
			RootMutationDescription,
			&allMutationFields,
		),
		newRootObject(
			RootSubscription,
			RootSubscriptionDescription,
			&allSubscriptionFields,
		),
		&allTypes,
		&allDirectives,
	)

	if err != nil {
		return nil, errors.Errorf("failed to create schema: %v", err)
	}

	return &schema, nil
}
