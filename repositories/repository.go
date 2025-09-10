package repositories

import (
	"user_management_service/model"

	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	CheckForExistingUser(field, value string) (*model.CommercialUser, error)
	CreateCommercialUser(signUpInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error)
	FetchUserByLoginID(field, value string) (*model.CommercialUser, error)
	CreateUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error)
	UpdateCommercialUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error)
	FetchProfileByUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error)
	UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.CommercialUser, error)
}
