package schema

import "github.com/graphql-go/graphql"

var CommercialAuthError = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommercialAuthError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})

var UserError = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
