package schema

import (
	"user_management_service/resolver"
	"user_management_service/graph/scalar"
	
	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.UserResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createCommercialUser": &graphql.Field{
				Type: CreateUserResponse,
				Args: graphql.FieldConfigArgument{
					"signup_input": &graphql.ArgumentConfig{
						Type: SignupInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(resolver.CreateCommercialUser)(p), nil
				},
			},
			"createUserProfile": &graphql.Field{
				Type: UserProfileResponse,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: UserProfileInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("create", resolver.CreateUserProfile))(p), nil
				},
			},
			"updateCommercialUser": &graphql.Field{
				Type: CreateUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: scalar.UUID,
					},
					"signup_input": &graphql.ArgumentConfig{
						Type: SignupInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("update", resolver.UpdateCommercialUser))(p), nil
				},
			},
			"updateCommercialUserStatus" : &graphql.Field{
				Type: UserStatusResponse,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("update", resolver.UpdateUserStatus))(p), nil
				},
			},
			"resetCommercialPassword": &graphql.Field{
				Type: GenericUserResponse,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(scalar.UUID),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"confirm_password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return AuthMiddleware(PermissionMiddleware("update", resolver.ResetPassword))(p), nil
				},
			},
		},
	})
}
