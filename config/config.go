package config

import (
    "os"
    "time"
)

func PostgresUser() string     { return os.Getenv("POSTGRES_USER") }
func PostgresPassword() string { return os.Getenv("POSTGRES_PASSWORD") }
func PostgresHost() string     { return os.Getenv("POSTGRES_HOST") }
func PostgresPort() string     { return os.Getenv("POSTGRES_PORT") }
func PostgresDB() string       { return os.Getenv("POSTGRES_DB") }

const (
    UserTable             = "users"
    AuthAccessTokenTable  = "access_tokens"
    AuthClientTable       = "auth_clients"
    AuthRefreshTokenTable = "refresh_tokens"
    AccessTokenDuration   = 3 * 30 * 24 * 60 * 60 * time.Second //3 months
    RefreshTokenDuration  = 6 * 30 * 24 * 60 * 60 * time.Second //6 months
    BcryptCost            = 12
)

func ClientId() string     { return os.Getenv("CLIENT_ID") }
func ClientSecret() string { return os.Getenv("CLIENT_SECRET") }
func RabbitMQURL() string { return os.Getenv("RABBITMQ_URL") }
func UserServiceApi() string { return os.Getenv("USER_SERVICE_API") }
func MembershipApi() string  { return os.Getenv("MEMBERSHIP_API") }
