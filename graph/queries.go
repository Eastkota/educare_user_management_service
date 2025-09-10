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
						Name:    "AuthService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
			"checkForExistingUser": &graphql.Field{
				Type: CheckForExistingUserResponse,
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
			"fetchProfileByUserId": &graphql.Field{
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
			// "fetchUser": &graphql.Field{
			// 	Type: SingleUserResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"user_id": &graphql.ArgumentConfig{
			// 			Type: scalar.UUID,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return AuthMiddleware(resolver.FetchUser)(p), nil
			// 	},
			// },
		},
	})
}
