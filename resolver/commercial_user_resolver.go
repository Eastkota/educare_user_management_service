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

    // Get the database connection from the context or a global variable.
    // Assuming 'ur.DB' holds the GORM database instance.
    db, err := helpers.GetGormDB()

    // Start a new transaction.
    tx := db.Begin()
    if tx.Error != nil {
        return helpers.FormatError(tx.Error)
    }

    // Defer a rollback in case of an error.
    // If the function returns successfully, the commit will prevent this.
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Pass the transaction 'tx' to the service layer.
    result, err := ur.Services.CreateUserProfile(userProfileInput, tx)
    if err != nil {
        tx.Rollback() // Roll back the transaction on error.
        return helpers.FormatError(err)
    }

    // Commit the transaction if everything was successful.
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
	// Extract the question ID and the new 'is_published' status from the GraphQL arguments.
	userID, ok := p.Args["userID"].(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return helpers.FormatError(fmt.Errorf("userID is required"))
	}

	status, ok := p.Args["status"].(string)
	if !ok {
		// A boolean argument is required, return an error if it's not present.
		return helpers.FormatError(fmt.Errorf("User status is required and must be a boolean"))
	}

	// Call the service layer to perform the status update.
	// Assuming the service layer has a corresponding method.
	result, err := ur.Services.UpdateUserStatus(p.Context, userID, status)
	if err != nil {
		return helpers.FormatError(err)
	}

	// Format the successful response.
	return &model.GenericUserResponse{
		Data: &model.DeleteUserResult{
			User: result,
		},
		Error: nil,
	}
}


// func (ar *UserResolver) UpdatePassword(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	var updatePasswordInput model.UpdatePasswordInput
// 	inputData := p.Args["input"].(map[string]interface{})

// 	jsonData, err := json.Marshal(inputData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = json.Unmarshal(jsonData, &updatePasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = ar.Services.UpdatePassword(updatePasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.GenericAuthSuccessData{
// 			Message: "Password updated successfully",
// 		},
// 		Error: nil,
// 	}
// }
// func (ar *UserResolver) UpdateSingleDataByID(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	// ctx := p.Context
//     // user := ctx.Value(model.UserKey).(*model.User)
// 	// if user == nil {
// 	// 	return helpers.FormatError(fmt.Errorf("unauthorized"))
// 	// }
		
// 	var updateSingleData model.UpdateSingleAuthDataInput
// 	inputData := p.Args["input"].(map[string]interface{})

// 	jsonData, err := json.Marshal(inputData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = json.Unmarshal(jsonData, &updateSingleData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	userID := uuid.MustParse("f81c3b55-d7a5-42ef-9f12-1d05dea507c6")//user.ID.String() /// TODO 
// 	//Implement  the login to fetch user form Authorization
// 	result, err := ar.Services.UpdateSingleDataByID(userID, updateSingleData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.UserResult{
// 			User: result,
// 		},
// 		Error: nil,
// 	}
// }

// func (ar *UserResolver) FetchUser(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	userID := p.Args["user_id"].(uuid.UUID)
// 	result, err := ar.Services.FetchUser(userID)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.UserResult{
// 			User: result,
// 		},
// 		Error: nil,
// 	}
// }

// func (ar *UserResolver) ResetPassword(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	var resetPasswordInput model.ResetPasswordInput
// 	inputData := p.Args["input"].(map[string]interface{})

// 	jsonData, err := json.Marshal(inputData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = json.Unmarshal(jsonData, &resetPasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = ar.Services.ResetPassword(resetPasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.GenericAuthSuccessData{
// 			Message: "Password updated successfully",
// 		},
// 		Error: nil,
// 	}
// }
