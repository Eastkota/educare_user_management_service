package model

import "github.com/google/uuid"

type SignupInput struct {
    Email    string `json:"email"`
    MobileNo string `json:"mobile_no"`
    Name     string `json:"name"`
    Gender   string `json:"gender"`
    Password string `json:"password"`
}

type UpdatePasswordInput struct {
    CurrentPassword string `json:"current_password"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
    UserId          uuid.UUID `json:"user_id"`
}

type UpdateSingleAuthDataInput struct {
    Field    string `json:"field"`
    Value    string `json:"value"`
    Password string `json:"password"`
}

type ResetPasswordInput struct {
    UserId          uuid.UUID `json:"user_id"`
    Password        string    `json:"password"`
    ConfirmPassword string    `json:"confirm_password"`
}

type UserProfileInput struct {
    Name           string `json:"name"`
    Gender         string `json:"gender"`
    ProfilePicture string `json:"profile_picture"`
    UserId         uuid.UUID `json:"user_id"`
}