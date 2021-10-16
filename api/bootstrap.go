package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/vaccine/api/route/user"
	"github.com/zakiafada32/vaccine/api/utils"
)

type Controller struct {
	User *user.UserController
}

func Bootstrap(e *echo.Echo, c *Controller) {
	if c.User == nil {
		panic("user controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userV1 := e.Group("api/vaccine-user")
	userV1.POST("/register", c.User.Create)
	userV1.GET("", c.User.FindAll)
}
