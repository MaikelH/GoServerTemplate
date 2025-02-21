package controllers

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"goservertemplate/httpserver/middleware"
	"goservertemplate/types"
	"net/http"
)

type UserController struct {
}

func (u *UserController) MountRoutes(s *fuego.Server, config *types.Configuration) {
	userGroup := fuego.Group(s, "user")

	fuego.Use(userGroup, middleware.EnsureValidToken(config))

	get := fuego.Get(userGroup, "/", func(context fuego.ContextNoBody) (any, error) {
		token := context.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

		claims := token.CustomClaims.(*middleware.CustomClaims)
		if !claims.HasScope("read:messages") {
			context.Response().WriteHeader(http.StatusForbidden)
			context.Response().Write([]byte(`{"message":"Insufficient scope."}`))
			return nil, nil
		}

		context.Response().WriteHeader(http.StatusOK)
		context.Response().Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		return "UserProfile", nil
	})
	get.Operation.Description = "Get user profile"
	get.Operation.Summary = "Get user profile"
	get.Operation.Security = &openapi3.SecurityRequirements{}
	get.Operation.Security.With(map[string][]string{"bearerAuth": {}})

}
