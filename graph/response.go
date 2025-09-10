package schema

import "github.com/graphql-go/graphql"


var GenericAuthResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericAuthResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: AuthGenericSuccessData},
		"error": &graphql.Field{Type: AuthError},
	},
})

var CheckForExistingUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckForExistingUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: ExistUser},
		"error": &graphql.Field{Type: AuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserResult},
		"error": &graphql.Field{Type: AuthError},
	},
})

var UserProfileResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserProfileResult},
		"error": &graphql.Field{Type: UserError},
	},
})

var GenericUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GenericUserSuccessData},
		"error": &graphql.Field{Type: UserError},
	},
})

var CreateUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CommercialUser},
		"error": &graphql.Field{Type: AuthError},
	},
})
var UserStatusResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserStatusResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CommercialUserStatus},
		"error": &graphql.Field{Type: AuthError},
	},
})

