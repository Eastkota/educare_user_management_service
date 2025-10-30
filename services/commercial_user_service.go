package services

import (
	"user_management_service/repositories"
	"user_management_service/model"

	"fmt"
	"context"
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserService struct {
	Repository repositories.Repository // Inject Repository
}

func NewUserService(repository repositories.Repository) *UserService {
	return &UserService{Repository: repository}
}

func (as *UserService) CheckForExistingUser(field, value string) (*model.CommercialUser, error) {
	return as.Repository.CheckForExistingUser(field, value)
}

func (as *UserService) CreateCommercialUser(signupData model.SignupInput) (*model.CommercialUser, *model.UserProfile, error) {

	if signupData.MobileNo == "" && signupData.Email == "" {
		return nil, nil, fmt.Errorf("either email or mobile number is required")
	}
	if signupData.MobileNo != "" {
		mobileResult, err := as.CheckForExistingUser("mobile_no", signupData.MobileNo)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register: %v", err)
		}

		if mobileResult != nil {
			return nil, nil, fmt.Errorf("user with given mobile number %v already exist", signupData.MobileNo)
		}
	}

	if signupData.Email != "" {
		emailResult, err := as.CheckForExistingUser("email", signupData.Email)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to register %v", err)
		}
		if emailResult != nil {
			return nil, nil, fmt.Errorf("user with given email address %v already exist", signupData.Email)
		}
	}

	return as.Repository.CreateCommercialUser(&signupData)
}

func (as *UserService) CreateUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error) {
    // Pass the transaction 'tx' to the repository method.
    return as.Repository.CreateUserProfile(tx, inputData)
}

func (as *UserService) UpdateCommercialUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error) {
	return as.Repository.UpdateCommercialUser(userID, signupInput)
}

func (as *UserService) FetchProfileByUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error) {
    return as.Repository.FetchProfileByUserId(ctx, userID)
}

func (vs *UserService) UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.CommercialUser, error) {
	return vs.Repository.UpdateUserStatus(ctx, userID, status)
}

func (vs *UserService) FetchCommercialUser(userID uuid.UUID) (*model.CommercialUser, error) {
	return vs.Repository.FetchCommercialUser(userID)
}

func (vs *UserService) FetchAllCommercialUsers(limit, offset int) ([]model.CommercialUser, int, error) {
    return vs.Repository.FetchAllCommercialUsers(limit, offset)
}

func (vs *UserService) FetchAllActiveUsers(limit, offset int) ([]model.CommercialUser, int, error) {
	return vs.Repository.FetchAllActiveUsers(limit, offset)
}
func (vs *UserService) FetchNewRegister(from_date, to_date time.Time) ([]model.CommercialUser, error) {
	return vs.Repository.FetchNewRegister(from_date, to_date)
}

func (vs *UserService) GetCommercialUserTotals(fromDate, toDate *time.Time) (totalAll int, totalActive int, totalNew int, err error) {
	return vs.Repository.GetCommercialUserTotals(fromDate, toDate)
}

func (vs *UserService) GetUserActivity(offset, limit int) ([]model.UserActivity, error) {
	return vs.Repository.GetUserActivity(offset, limit)
}

func (vs *UserService) ResetPassword(userID uuid.UUID, password, confirmPassword string) error {
    if password == "" || confirmPassword == "" {
        return fmt.Errorf("new password and confirm password cannot be empty")
    }
    if password != confirmPassword {
        return fmt.Errorf("new password and confirm password do not match")
    }

    // Call the new, dedicated repository function.
    // The repository handles the hashing and database logic.
    err := vs.Repository.ResetPassword(userID, password)
    if err != nil {
        return fmt.Errorf("failed to reset password: %v", err)
    }

    return nil
}


