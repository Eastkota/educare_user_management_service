package services

import (
	"user_management_service/model"

	"context"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Services interface {
	CheckForExistingUser(field, value string) (*model.CommercialUser, error)
	CreateCommercialUser(signupData model.SignupInput) (*model.CommercialUser, *model.UserProfile, error)
	CreateUserProfile(inputData model.UserProfileInput, tx *gorm.DB) (*model.UserProfile, error)
	UpdateCommercialUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error)
	FetchProfileByUserId(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
	UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.CommercialUser, error)
}