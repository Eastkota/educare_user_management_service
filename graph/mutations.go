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
					return resolver.CreateCommercialUser(p), nil
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
					return resolver.CreateUserProfile(p), nil
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
					return resolver.UpdateCommercialUser(p), nil
				},
			},
			"updateUserStatus" : &graphql.Field{
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
					return resolver.UpdateUserStatus(p), nil
				},
			},
			// "updatePassword": &graphql.Field{
			// 	Type: GenericAuthResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: UpdatePasswordInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.UpdatePassword(p), nil
			// 	},
			// },
			// "resetPassword": &graphql.Field{
			// 	Type: GenericAuthResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: ResetPasswordInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.ResetPassword(p), nil
			// 	},
			// },
			// "updateSingleUserDataById": &graphql.Field{
			// 	Type: SingleUserResponse,
			// 	Args: graphql.FieldConfigArgument{
			// 		"input": &graphql.ArgumentConfig{
			// 			Type: UpdateSingleAuthDataInput,
			// 		},
			// 	},
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return resolver.UpdateSingleDataByID(p), nil
			// 	},
			// },
		},
	})
}
