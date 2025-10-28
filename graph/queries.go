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
						Name:    "usermanagementservice",
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
					return AuthMiddleware(PermissionMiddleware("list", resolver.CheckForExistingUser))(p), nil
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
					return AuthMiddleware(PermissionMiddleware("list", resolver.FetchProfileByUserId))(p), nil
				},
			},
			"fetchCommercialUser": &graphql.Field{
				Type: SingleCommercialUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("list", resolver.FetchCommercialUser))(p), nil
				},
			},
			"fetchAllCommercialUsers": &graphql.Field{
				Type: UsersResponse,
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
						Description: "Maximum number of users to fetch per request.",
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
						Description: "The number of users to skip before starting to return results.",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("list", resolver.FetchAllCommercialUsers))(p), nil
				},
			},
			"fetchAllActiveCommercialUsers": &graphql.Field{
				Type: UsersResponse,
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
						Description: "Maximum number of users to fetch per request.",
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
						Description: "The number of users to skip before starting to return results.",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("list", resolver.FetchAllActiveUsers))(p), nil
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
					return AuthMiddleware(PermissionMiddleware("list", resolver.FetchNewRegister))(p), nil
				},
			},
			"getCommercialUserTotals": &graphql.Field{
                Type: UserTotalsResponse,
                Args: graphql.FieldConfigArgument{
                    "from_date": &graphql.ArgumentConfig{
                        Type: scalar.Time,
                    },
                    "to_date": &graphql.ArgumentConfig{
                        Type: scalar.Time,
                    },
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    return AuthMiddleware(PermissionMiddleware("list", resolver.GetCommercialUserTotals))(p), nil
                },
            },
		},
	})
}
