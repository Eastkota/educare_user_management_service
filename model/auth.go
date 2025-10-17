package model
import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name           string    `gorm:"type:varchar(50)" json:"name"`
    Identifier     string    `gorm:"type:varchar(255)" json:"identifier"`
    Email          string    `gorm:"type:varchar(255);unique" json:"email"`
    MobileNo       string    `gorm:"type:varchar(20);unique" json:"mobile_no"`
    Password       string    `gorm:"type:text;not null" json:"password_hash"`
    Status         string    `gorm:"type:varchar(50);default:active;not null" json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`

    Roles          []Role    `gorm:"many2many:cms_auth.user_roles;" json:"roles"`
}

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`

    Users       []User `gorm:"many2many:user_roles;" json:"users"`
    Permissions []Permission `gorm:"foreignKey:RoleId" json:"permissions"`
}

type Permission struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    RoleId      uuid.UUID `gorm:"type:uuid" json:"role_id"`
    ActionId    uuid.UUID `gorm:"type:uuid" json:"action_id"`

    Role        *Role     `gorm:"foreignKey:RoleId" json:"role"`
    Action      *Action   `gorm:"foreignKey:ActionId" json:"action"`

    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Action struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string `gorm:"type:varchar(50)" json:"name"`
    Action      string `gorm:"type:varchar(50)" json:"action"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`

    ResourceID  uuid.UUID   `gorm:"type:uuid" json:"resource_id"`
    Resource    *Resource   `gorm:"foreignKey:ResourceID"`
}

type Resource struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string `gorm:"type:varchar(50)" json:"name"`
    Menu      string `gorm:"type:varchar(50)" json:"menu"`
    Display      string `gorm:"type:varchar(50)" json:"display"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`

    ModuleID  uuid.UUID   `gorm:"type:uuid" json:"module_id"`
    Module    *Module   `gorm:"foreignKey:ModuleID"`
}

type Module struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string         `gorm:"type:varchar(255)" json:"name"`
    Menu       string         `gorm:"type:text" json:"menu"`
    Display       string         `gorm:"type:text" json:"display"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
}