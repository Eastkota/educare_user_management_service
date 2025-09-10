package handlers

import (
	"user_management_service/graph"
	
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

func Handler(ctx echo.Context) error {
	if ctx.Request().Method != http.MethodPost {
		return echo.ErrMethodNotAllowed
	}

	var operation map[string]interface{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&operation)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query, ok := operation["query"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing query field")
	}

	var variables map[string]interface{}
	if opVars, ok := operation["variables"]; ok && opVars != nil {
		variables = opVars.(map[string]interface{})
	} else {
		variables = make(map[string]interface{})
	}

	result := executeQuery(ctx.Request().Context(), query, variables)

	if result == nil {
        // Handle the error appropriately, e.g., return a 500 status code.
        return ctx.JSON(http.StatusInternalServerError, map[string]string{
            "error": "GraphQL query execution failed",
        })
    }

	if result.HasErrors() {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Errors)
	}

	return ctx.JSON(http.StatusOK, result)
}

func executeQuery(ctx context.Context, query string, variables map[string]interface{}) *graphql.Result {
	authSchema, err := schema.GetSchema()
	if err != nil {
		return nil
	}
	res := graphql.Do(graphql.Params{
		Schema:         *authSchema,
		RequestString:  query,
		VariableValues: variables,
		Context:        ctx,
	})
	return res
}
