package repositories

import (
    "user_management_service/model"
    "user_management_service/helpers"

    "fmt"
    "time"
    "errors"
    "context"

    "gorm.io/gorm"
	"github.com/google/uuid"
)

type UserRepository struct{
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (repo *UserRepository) CheckForExistingUser(field, value string) (*model.CommercialUser, error) {
    var user model.CommercialUser
    err := repo.DB.Where(fmt.Sprintf("%s = ? AND status != ?", field), value, "Deleted").First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &user, nil
}

func (repo *UserRepository) FetchUserByLoginID(field, value string) (*model.CommercialUser, error) {
    var user model.CommercialUser
    err := repo.DB.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &user, nil
}


// CreateCommercialUser creates a new commercial user and their associated profile within a single transaction.
func (repo *UserRepository) CreateCommercialUser(signupInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error) {
    var user *model.CommercialUser
    var userProfile *model.UserProfile
    var err error

    err = repo.DB.Transaction(func(tx *gorm.DB) error {
        identifier, err := helpers.GenerateRandomTokenString(6)
        if err != nil {
            return fmt.Errorf("failed to generate identifier: %v", err)
        }

        hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
        if err != nil {
            return fmt.Errorf("failed to hash password: %v", err)
        }

        // First, check for an existing user by mobile number or email.
        // NOTE: This check should also be part of the transaction to prevent race conditions.
        // The current implementation is simplified for this example.
        if signupInput.MobileNo != "" {
            user, _ = repo.FetchUserByLoginID("mobile_no", signupInput.MobileNo)
            if user == nil && signupInput.Email != "" {
                user, _ = repo.FetchUserByLoginID("email", signupInput.Email)
            }
        }

        // If a user is found, update their record.
        if user != nil {
            updateData := map[string]interface{}{
                "mobile_no":       signupInput.MobileNo,
                "email":           signupInput.Email,
                "user_identifier": identifier,
                "password":        hashedPassword,
                "status":          "Active",
                "updated_at":      time.Now(),
            }
            if err := tx.Model(user).Updates(updateData).Error; err != nil {
                return fmt.Errorf("failed to update user data: %v", err)
            }
            // Note: If you want to return the user's profile on update,
            // you'll need to fetch it here.
            return nil
        }

        // If no user is found, create a new user and their profile.
        newUser := model.CommercialUser{
            ID:             uuid.New(),
            Name:           signupInput.Name,
            MobileNo:       signupInput.MobileNo,
            Email:          signupInput.Email,
            UserIdentifier: identifier,
            Password:       hashedPassword,
            Status:         "Active",
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }

        if err := tx.Create(&newUser).Error; err != nil {
            return fmt.Errorf("failed to insert user data: %v", err)
        }
        user = &newUser

        profileInput := model.UserProfileInput{
            UserId: newUser.ID,
            Gender: signupInput.Gender,
            Name:   signupInput.Name,
        }
        // Pass the transaction 'tx' to the CreateUserProfile function.
        userProfile, err = repo.CreateUserProfile(tx, profileInput)
        if err != nil {
            return fmt.Errorf("failed to create user profile: %v", err)
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }

    return user, userProfile, nil
}

func (repo *UserRepository) CreateUserProfile(tx *gorm.DB, inputData model.UserProfileInput) (*model.UserProfile, error) {
    userProfile := model.UserProfile{
        ID:             uuid.New(),
        Name:           inputData.Name,
        Gender:         inputData.Gender,
        ProfilePicture: inputData.ProfilePicture,
        UserId:         inputData.UserId,
    }

    // Use the passed-in transaction object for all database operations.
    if err := tx.Create(&userProfile).Error; err != nil {
        return nil, fmt.Errorf("failed to insert user profile: %v", err)
    }

    favoritePlaylist := model.UserVideoPlaylist{
        ID:        uuid.New(),
        Name:      "Favorites",
        Favorites: true,
        ProfileId: userProfile.ID,
    }

    if err := tx.Create(&favoritePlaylist).Error; err != nil {
        return nil, fmt.Errorf("failed to create favorite playlist: %v", err)
    }

    userProfile.FavoriteVideoPlaylistId = favoritePlaylist.ID
    // Use the passed-in transaction object.
    if err := tx.Save(&userProfile).Error; err != nil {
        return nil, fmt.Errorf("failed to update user profile with playlist ID: %v", err)
    }

    return &userProfile, nil
}

func (repo *UserRepository) FetchProfileByUserId(ctx context.Context, userId uuid.UUID) (*model.UserProfile, error) {
    var profile model.UserProfile

    db, err := helpers.GetGormDB()
    if err != nil {
        return nil, err
    }

    if err := db.WithContext(ctx).Where("user_id = ?", userId).First(&profile).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch profile for user id '%s': %w", userId, err)
    }
    return &profile, nil
}

func (repo *UserRepository) UpdateCommercialUser(userID uuid.UUID, signupInput *model.SignupInput) (*model.CommercialUser, *model.UserProfile, error) {
    var user model.CommercialUser
    var userProfile *model.UserProfile
    var err error

    err = repo.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.First(&user, "id = ?", userID).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return fmt.Errorf("user with ID %s not found: %w", userID, err)
            }
            return fmt.Errorf("failed to fetch user for update: %w", err)
        }

        // Step 2: Prepare the update data.
        hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
        if err != nil {
            return fmt.Errorf("failed to hash password: %w", err)
        }

        updateData := map[string]interface{}{
            "name":       signupInput.Name,
            "password":   hashedPassword,
            "updated_at": time.Now(),
        }
        
        if signupInput.MobileNo != "" {
            updateData["mobile_no"] = signupInput.MobileNo
        }
        if signupInput.Email != "" {
            updateData["email"] = signupInput.Email
        }

        if err := tx.Model(&model.CommercialUser{}).Where("id = ?", userID).Updates(updateData).Error; err != nil {
            return fmt.Errorf("failed to update user data: %w", err)
        }

        // The 'user' variable is from the tx.First call, so we can use its ID.
        userProfile, err = repo.FetchProfileByUserId(context.Background(), user.ID)
        if err != nil {
            return fmt.Errorf("failed to fetch user profile after update: %w", err)
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }


    var updatedUser model.CommercialUser
    if err := repo.DB.First(&updatedUser, "id = ?", userID).Error; err != nil {
        return nil, nil, fmt.Errorf("failed to re-fetch updated user: %w", err)
    }

    return &updatedUser, userProfile, nil
}

func (repo *UserRepository) UpdateUserStatus(ctx context.Context, userID uuid.UUID, status string) (*model.CommercialUser, error) {
	// Find the existing question by its ID
	var user model.CommercialUser
	if err := repo.DB.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("user not found with ID: %s", userID)
	}

	updates := map[string]interface{}{
		"status": status,
		"updated_at":   time.Now(),
	}

	if err := repo.DB.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update user status: %w", err)
	}

	// Fetch the updated question to return the new state.
	var updatedUser model.CommercialUser
	if err := repo.DB.WithContext(ctx).Where("id = ?", userID).First(&updatedUser).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch the updated USER: %w", err)
	}

	return &updatedUser, nil
}


