package schema

import (
	"user_management_service/graph/scalar"
	
	"github.com/graphql-go/graphql"
)

var UserMembership = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUserMembership",
	Fields: graphql.Fields{
		"id":                     &graphql.Field{Type: scalar.UUID},
		"membership_join_date":   &graphql.Field{Type: scalar.Time},
		"membership_end_date":    &graphql.Field{Type: scalar.Time},
		"user_id":             &graphql.Field{Type: scalar.UUID},
		"membership_duration_id": &graphql.Field{Type: scalar.UUID},
	},
})
