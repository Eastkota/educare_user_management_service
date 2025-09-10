package model

type GenericAuthSuccessData struct {
    Message string `json:"message"`
    Code    string `json:"code"`
}

type GenericAuthResponse struct {
    Data  interface{} `json:"data,omitempty"`
    Error *AuthError  `json:"error,omitempty"`
}

type CreateUserSuccessData struct {
    User       *CommercialUser              `json:"user"`
    Profile    *UserProfile   `json:"profile"`
}
type DeleteUserResult struct {
    User       *CommercialUser              `json:"user"`
}

type GenericUserResponse struct {
	Data  interface{}
	Error *UserError
}

type UserProfileResult struct {
	UserProfile *UserProfile `json:"user_profile"`
}