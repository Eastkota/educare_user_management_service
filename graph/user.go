package schema

import (
	"user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var CmsUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "CmsUser",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: scalar.UUID},
		"user_identifier": &graphql.Field{Type: graphql.String},
		"email":           &graphql.Field{Type: graphql.String},
		"mobile_no":       &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"created_at":      &graphql.Field{Type: scalar.Time},
		"updated_at":      &graphql.Field{Type: scalar.Time},

		"user_profile":	   &graphql.Field{Type: CmsUserProfile},
	},
})

// Define UserType
var CmsUserProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "CmsUserProfile",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"gender":                     &graphql.Field{Type: graphql.String},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})


var CmsExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "CmsExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})
var CmsUserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "CmsUserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: CmsUser},
	},
})

var CmsUserProfileResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "CmsUserProfileResult",
	Fields: graphql.Fields{
		"user_profile": &graphql.Field{Type: CmsUserProfile},
	},
})

var AuthGenericSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthGenericSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var GenericUserSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericUserSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var CommercialUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommercialUser",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: CmsUser},
		"profile":    &graphql.Field{Type: CmsUserProfile},
	},
})
var CommercialUserStatus = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommercialUserStatus",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: CmsUser},
	},
})

var FetchAllUsersResult = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "FetchAllUsersResult",
        Fields: graphql.Fields{
            "users": &graphql.Field{
                Type: graphql.NewList(CmsUser),
            },
        },
    },
)
