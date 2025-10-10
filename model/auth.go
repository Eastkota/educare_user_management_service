package model
import (
	"time"
	"github.com/google/uuid"
)

type User struct {
    ID             uuid.UUID `json:"id"`
    UserIdentifier string    `json:"user_identifier"`
    Email          string    `json:"email"`
    MobileNo       string    `json:"mobile_no"`
    Password       string    `json:"password_hash"`
    Status         string    `json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}