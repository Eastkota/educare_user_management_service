package schema

import (
    "user_management_service/helpers"
    "user_management_service/model"
    
    "fmt"
    "context"
    "net/http"

    "github.com/graphql-go/graphql"
)

func AuthMiddleware(next func(p graphql.ResolveParams) *model.GenericUserResponse) func(p graphql.ResolveParams) *model.GenericUserResponse {
    return func(p graphql.ResolveParams) *model.GenericUserResponse {
        ctx := p.Context
        userInterface := ctx.Value("user")

        var user *model.User
        if userInterface == nil {
            if req, ok := ctx.Value("http_request").(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := helpers.ValidateToken(authHeader)
                if err != nil {
                    return helpers.FormatError(err)
                }
                if u == nil {
                    return helpers.FormatError(fmt.Errorf("invalid_token"))
                }

                ctx = context.WithValue(ctx, "user", u)
                p.Context = ctx
                user = u
            } else {
                return helpers.FormatError(fmt.Errorf("invalid_token"))
            }
        } else {
            user, _ = userInterface.(*model.User)
        }

        if user == nil {
            return helpers.FormatError(fmt.Errorf("invalid_token"))
        }
        return next(p)
    }
}

func PermissionMiddleware(actionName string, next func(p graphql.ResolveParams) *model.GenericUserResponse) func(p graphql.ResolveParams) *model.GenericUserResponse {
    return func(p graphql.ResolveParams) *model.GenericUserResponse {
        ctx := p.Context
        user := ctx.Value("user")
        if user == nil {
            return helpers.FormatError(fmt.Errorf("UnAuthorized"))
        }

        userData := user.(*model.User)
        hasPermission := false

        for _, role := range userData.Roles {
            if role.Name == "super admin" {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            for _, role := range userData.Roles {
                for _, permission := range role.Permissions {
                    if permission.Action.Action == actionName {
                        hasPermission = true
                        break 
                    }
                }
                if hasPermission {
                    break
                }
            }
        }

        if !hasPermission {
            fmt.Println("User does not have permission for action:", actionName)
            return helpers.FormatError(fmt.Errorf("UnAuthorized"))
        }

        return next(p)
    }
}