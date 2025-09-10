package model

type AuthError struct {
    Message string `json:"message"`
    Code    string `json:"code"`
    Field   string `json:"field,omitempty"`
}

type UserError struct {
    Message string `json:"message"`
    Code    string `json:"code"`
    Field   string `json:"field,omitempty"`
}