package controllers

import "github.com/go-fuego/fuego"

type UserController struct {
}

func (u *UserController) MountRoutes(s *fuego.Server) {
	userGroup := fuego.Group(s, "user")

	get := fuego.Get(userGroup, "/", func(context fuego.ContextNoBody) (any, error) {
		return "UserProfile", nil
	})
	get.Operation.Description = "Get user profile"
	get.Operation.Summary = "Get user profile"

}
