package model

import (
    "time"

    "github.com/google/uuid"
)

type ContextKey string

const (
    UserKey    ContextKey = "user"
    RequestKey ContextKey = "http_request"
)

type CommercialUser struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name           string    `gorm:"type:varchar(100);not null" json:"name"`
    UserIdentifier string    `gorm:"type:varchar(32);unique;not null" json:"user_identifier"`
    Email          string    `gorm:"type:varchar(100);unique" json:"email"`
    MobileNo       string    `gorm:"type:varchar(20);unique" json:"mobile_no"`
    Password       string    `gorm:"type:text;not null" json:"password_hash"`
    Status         string    `gorm:"type:varchar(50);default:active;not null" json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (CommercialUser) TableName() string {
    return "auth.users"
}

type UserResult struct {
    CommercialUser *CommercialUser `json:"user"`
}

type UserProfile struct {
    ID                      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name                    string    `gorm:"type:varchar(100)" json:"name"`
    ProfilePicture          string    `gorm:"type:varchar(255)" json:"profile_picture"`
    Gender                  string    `gorm:"type:char(1)" json:"gender"`
    UserId                  uuid.UUID `gorm:"type:uuid" json:"user_id"`
    FavoriteVideoPlaylistId uuid.UUID    `gorm:"type:uuid" json:"favorite_video_playlist_id"`
    CreatedAt               time.Time `json:"created_at"`
    UpdatedAt               time.Time `json:"updated_at"`
}

func (UserProfile) TableName() string {
    return "user_data.user_profiles"
}

type UserVideoPlaylist struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(100)" json:"name"`
    Favorites bool      `gorm:"type:boolean;default:false" json:"favorites"`
    ProfileId uuid.UUID `gorm:"type:uuid" json:"profile_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (UserVideoPlaylist) TableName() string {
    return "user_data.video_playlists"
}
