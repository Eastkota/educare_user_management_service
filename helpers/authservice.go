package helpers

import (
	"user_management_service/config"
	"user_management_service/model"

	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func ValidateToken(tokenStr string) (*model.User, error) {
	authServiceClient := graphql.NewClient(config.AuthServiceApi())
	req := graphql.NewRequest(`
		query ValidateToken($input:  String){
			validateToken(token: $input) {
				data {
					user {
						created_at
						email
						id
						mobile_no
						name
						status
						updated_at
						user_identifier
						roles {
							description
							id
							name
							permissions {
								created_at
								id
								role_id
								updated_at
								action {
									action
									created_at
									id
									name
									updated_at
									resource {
										created_at
										display
										id
										menu
										name
										updated_at
										module {
											created_at
											display
											id
											menu
											name
											updated_at
										}
									}
								}
							}
						}
					}
				}
				error {
					code
					field
					message
				}
			}
		}
	`)
	req.Var("input", tokenStr)
	req.Header.Set("Cache-Control", "no-cache")

	var response struct {
		ValidateToken struct {
			Data struct {
				User model.User `json:"user"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"validateToken"`
	}

	err := authServiceClient.Run(context.Background(), req, &response)
	if err != nil {
		return nil, fmt.Errorf("invalid_token: %v", err)
	}
	if response.ValidateToken.Error.Message != "" {
		return nil, fmt.Errorf(response.ValidateToken.Error.Message)
	}
	return &response.ValidateToken.Data.User, err
}