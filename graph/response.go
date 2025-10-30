package schema

import "github.com/graphql-go/graphql"


var checkForExistingCommercialUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "checkForExistingCommercialUser",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: CmsExistUser},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})

var SingleCommercialUserResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "SingleCommercialUserResponse",
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
	Name: "CreateUserResponse",
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

var UserTotalsResponse = graphql.NewObject(graphql.ObjectConfig{
    Name: "UserTotalsResponse",
    Fields: graphql.Fields{
        "data": &graphql.Field{
            Type: graphql.NewObject(graphql.ObjectConfig{
                Name: "UserTotalsData",
                Fields: graphql.Fields{
                    "total_all":    &graphql.Field{Type: graphql.Int},
                    "total_active": &graphql.Field{Type: graphql.Int},
                    "total_new":    &graphql.Field{Type: graphql.Int},
                },
            }),
        },
        "error": &graphql.Field{Type: graphql.String},
    },
})

var UserActivityResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserActivityResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: UserActivityResultType},
		"error": &graphql.Field{Type: CommercialAuthError},
	},
})

var UserActivityResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserActivityResultType",
	Fields: graphql.Fields{
		"user_activity": &graphql.Field{Type: graphql.NewList(UserActivity)},
	},
})


