package schema

import (
	"user_management_service/helpers"
	"user_management_service/model"
	"user_management_service/resolver"
	"user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

func NewQueryType(resolver *resolvers.UserResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
					if err != nil {
						return nil, err
					}

					serviceInfo := model.Service{
						Name:    "UserManagementService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
			"checkForExistingCommercialUser": &graphql.Field{
				Type: checkForExistingCommercialUser,
				Args: graphql.FieldConfigArgument{
					"field": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"value": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.CheckForExistingUser(p), nil
				},
			},
			"fetchProfileByCommercialUserId": &graphql.Field{
				Type: UserProfileResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchProfileByUserId(p), nil
				},
			},
			"fetchUser": &graphql.Field{
				Type: SingleUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchUser(p), nil
				},
			},
			"fetchAllCommercialUsers": &graphql.Field{
				Type: UsersResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllUsers(p), nil
				},
			},
			"fetchAllActiveCommercialUsers": &graphql.Field{
				Type: UsersResponse,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchAllActiveUsers(p), nil
				},
			},
			"fetchCommercialNewRegister": &graphql.Field{
				Type: UsersResponse,
				Args: graphql.FieldConfigArgument{
					"from_date": &graphql.ArgumentConfig{
						Type: scalar.Time,
					},
					"to_date": &graphql.ArgumentConfig{
						Type: scalar.Time,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.FetchNewRegister(p), nil
				},
			},
		},
	})
}
