package user

import (
	"github.com/labstack/echo/v4"

	message "github.com/zakiafada32/vaccination-record/business"
	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	"github.com/zakiafada32/vaccination-record/utils"
)

type UserController struct {
	service userBusiness.Service
}

func NewUserController(service userBusiness.Service) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var body registerUserRequestBody
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(utils.ConstructResponse(message.BadRequest, echo.Map{}))
	}

	err = c.Validate(&body)
	if err != nil {
		return err
	}

	id, err := uc.service.RegisterUser(body.convertToUserBusiness())
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.SuccessCreated, echo.Map{
		"id": id,
	}))
}

func (uc *UserController) FindAll(c echo.Context) error {
	users, err := uc.service.FindAllUser()
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.Success, echo.Map{
		"users": users,
	}))
}

func (uc *UserController) CheckStatus(c echo.Context) error {
	var body checkStatusRequestBody
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(utils.ConstructResponse(message.BadRequest, echo.Map{}))
	}

	userStatus, err := uc.service.CheckStatus(body.IdentityCardNumber, body.Name, body.Password)
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.Success, echo.Map{
		"status": userStatus,
	}))
}
