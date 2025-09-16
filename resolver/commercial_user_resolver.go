package resolvers

import (
	"user_management_service/helpers"
	"user_management_service/model"
	"user_management_service/services"

	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type UserResolver struct {
	Services services.Services // Inject Services
}

func NewUserResolver(service services.Services) *UserResolver {
	return &UserResolver{Services: service}
}

func (ar *UserResolver) CheckForExistingUser(p graphql.ResolveParams) *model.GenericUserResponse {
	field := p.Args["field"].(string)
	value := p.Args["value"].(string)
	result, err := ar.Services.CheckForExistingUser(field, value)
	if err != nil {
		return helpers.FormatError(err)
	}

	if result == nil {
        return &model.GenericUserResponse{
            Data: map[string]interface{}{
                "exist_user": false,
                "user_id":    nil,
            },
            Error: nil,
        }
    }
	
	return &model.GenericUserResponse{
		Data: map[string]interface{}{
			"exist_user": result != nil,
			"user_id":    result.ID,
		},
		Error: nil,
	}

}

func (ar *UserResolver) CreateCommercialUser(p graphql.ResolveParams) *model.GenericUserResponse {

	var signupInput model.SignupInput
	inputData := p.Args["signup_input"].(map[string]interface{})

	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = json.Unmarshal(jsonData, &signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	user, profile, err := ar.Services.CreateCommercialUser(signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericUserResponse{
		Data: &model.CreateUserSuccessData{
			User:    user,
			Profile: profile,
		},
		Error: nil,
	}
}

func (ur *UserResolver) CreateUserProfile(p graphql.ResolveParams) *model.GenericUserResponse {
    var userProfileInput model.UserProfileInput
    inputData := p.Args["input"].(map[string]interface{})
    jsonData, err := json.Marshal(inputData)
    if err != nil {
        return helpers.FormatError(err)
    }
    err = json.Unmarshal(jsonData, &userProfileInput)
    if err != nil {
        return helpers.FormatError(err)
    }
    db, err := helpers.GetGormDB()

    tx := db.Begin()
    if tx.Error != nil {
        return helpers.FormatError(tx.Error)
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    result, err := ur.Services.CreateUserProfile(userProfileInput, tx)
    if err != nil {
        tx.Rollback()
        return helpers.FormatError(err)
    }

    if err := tx.Commit().Error; err != nil {
        return helpers.FormatError(err)
    }

    return &model.GenericUserResponse{
        Data: &model.UserProfileResult{
            UserProfile: result,
        },
        Error: nil,
    }
}

func (ur *UserResolver) FetchProfileByUserId(p graphql.ResolveParams) *model.GenericUserResponse {
	userID := p.Args["user_id"].(uuid.UUID)
	result, err := ur.Services.FetchProfileByUserId(p.Context, userID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericUserResponse{
		Data: &model.UserProfileResult{
			UserProfile: result,
		},
		Error: nil,
	}
}

func (ur *UserResolver) UpdateCommercialUser(p graphql.ResolveParams) *model.GenericUserResponse {
    var signupInput model.SignupInput
    userID, ok := p.Args["user_id"].(uuid.UUID)
    if !ok {
        return helpers.FormatError(fmt.Errorf("user_id argument is not a valid UUID type"))
    }

    inputData, ok := p.Args["signup_input"].(map[string]interface{})
    if !ok {
        return helpers.FormatError(fmt.Errorf("signup_input argument is not a valid map"))
    }

    jsonData, err := json.Marshal(inputData)
    if err != nil {
        return helpers.FormatError(err)
    }

    err = json.Unmarshal(jsonData, &signupInput)
    if err != nil {
        return helpers.FormatError(err)
    }

    user, profile, err := ur.Services.UpdateCommercialUser(userID, &signupInput)
    if err != nil {
        return helpers.FormatError(err)
    }

    return &model.GenericUserResponse{
        Data: &model.CreateUserSuccessData{
            User:    user,
            Profile: profile,
        },
        Error: nil,
    }
}

func (ur *UserResolver) UpdateUserStatus(p graphql.ResolveParams) *model.GenericUserResponse {
	userID, ok := p.Args["userID"].(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return helpers.FormatError(fmt.Errorf("userID is required"))
	}

	status, ok := p.Args["status"].(string)
	if !ok {
		return helpers.FormatError(fmt.Errorf("User status is required and must be a boolean"))
	}

	result, err := ur.Services.UpdateUserStatus(p.Context, userID, status)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericUserResponse{
		Data: &model.DeleteUserResult{
			User: result,
		},
		Error: nil,
	}
}

func (ar *UserResolver) FetchUser(p graphql.ResolveParams) *model.GenericUserResponse {
	userID := p.Args["user_id"].(uuid.UUID)
	result, err := ar.Services.FetchUser(userID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericUserResponse{
		Data: &model.FetchSingleUserResult{
			User: result,
		},
		Error: nil,
	}
}

func (ar *UserResolver) FetchAllUsers(p graphql.ResolveParams) *model.GenericUserResponse {
    result, err := ar.Services.FetchAllUsers()
    if err != nil {
        return helpers.FormatError(err)
    }
    return &model.GenericUserResponse{
        Data: &model.FetchAllUsersResult{
            Users: result,
        },
        Error: nil,
    }
}

func (ar *UserResolver) FetchAllActiveUsers(p graphql.ResolveParams) *model.GenericUserResponse {
    result, err := ar.Services.FetchAllActiveUsers()
    if err != nil {
        return helpers.FormatError(err)
    }
    return &model.GenericUserResponse{
        Data: &model.FetchAllActiveUsersResult{
            Users: result,
        },
        Error: nil,
    }
}

func (ar *UserResolver) ResetPassword(p graphql.ResolveParams) *model.GenericUserResponse {
	userID := p.Args["user_id"].(uuid.UUID)
	password := p.Args["password"].(string)
	confirmPassword := p.Args["confirm_password"].(string)
	err := ar.Services.ResetPassword(userID, password, confirmPassword)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericUserResponse{
		Data: &model.GenericAuthSuccessData{
			Message: "Password updated successfully",
		},
		Error: nil,
	}
}


