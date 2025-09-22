package schema

import "github.com/graphql-go/graphql"


var checkForExistingCommercialUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "checkForExistingCommercialUser",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CmsExistUser},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})

var SingleUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleUserResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CmsUserResult},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})
var UsersResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UsersResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: FetchAllUsersResult},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})

var UserProfileResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CmsUserProfileResult},
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
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})
var UserStatusResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserStatusResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CommercialUserStatus},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})


