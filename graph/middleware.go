package schema

import (
	"user_management_service/helpers"
	"user_management_service/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

func AuthMiddleware(next func(p graphql.ResolveParams) *model.GenericUserResponse) func(p graphql.ResolveParams) *model.GenericUserResponse {
	return func(p graphql.ResolveParams) *model.GenericUserResponse {
		ctx := p.Context
		user := ctx.Value(model.UserKey).(*model.CommercialUser)
		if user == nil {
			return helpers.FormatError(fmt.Errorf("UnAuthorized"))
		}
		return next(p)
	}
}
