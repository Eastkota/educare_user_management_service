package schema

import (
	"user_management_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

// Define UserType
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: scalar.UUID},
		"user_identifier": &graphql.Field{Type: graphql.String},
		"name":            &graphql.Field{Type: graphql.String},
		"email":           &graphql.Field{Type: graphql.String},
		"mobile_no":       &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"created_at":      &graphql.Field{Type: scalar.Time},
		"updated_at":      &graphql.Field{Type: scalar.Time},
	},
})

// Define UserType
var UserProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUserProfile",
	Fields: graphql.Fields{
		"id":                         &graphql.Field{Type: scalar.UUID},
		"name":                       &graphql.Field{Type: graphql.String},
		"profile_picture":            &graphql.Field{Type: graphql.String},
		"favorite_video_playlist_id": &graphql.Field{Type: scalar.UUID},
		"gender":                     &graphql.Field{Type: graphql.String},
		"created_at":                 &graphql.Field{Type: scalar.Time},
		"updated_at":                 &graphql.Field{Type: scalar.Time},
	},
})


var ExistUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExistUser",
	Fields: graphql.Fields{
		"exist_user": &graphql.Field{Type: graphql.Boolean},
		"user_id":    &graphql.Field{Type: scalar.UUID},
	},
})
var UserResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResult",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: User},
	},
})

var UserProfileResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResult",
	Fields: graphql.Fields{
		"user_profile": &graphql.Field{Type: UserProfile},
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
		"user":       &graphql.Field{Type: User},
		"profile":    &graphql.Field{Type: UserProfile},
	},
})
var CommercialUserStatus = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommercialUserStatus",
	Fields: graphql.Fields{
		"user":       &graphql.Field{Type: User},
	},
})
