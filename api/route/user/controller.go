package user

import (
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/vaccine/api/utils"
	message "github.com/zakiafada32/vaccine/business"
	userBusiness "github.com/zakiafada32/vaccine/business/user"
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

	id, err := uc.service.Create(body.convertToUserBusiness())
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.SuccessCreated, echo.Map{
		"id": id,
	}))
}

func (uc *UserController) FindAll(c echo.Context) error {
	users, err := uc.service.FindAll()
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	usersData := make([]userResponse, len(users))
	for i, user := range users {
		usersData[i] = convertToUserResponse(user)
	}

	return c.JSON(utils.ConstructResponse(message.SuccessCreated, echo.Map{
		"users": usersData,
	}))
}

func (uc *UserController) AddHistory(c echo.Context) error {
	return nil
}

func (uc *UserController) DeleteHistory(c echo.Context) error {
	return nil
}

func (uc *UserController) FindByIdCard(c echo.Context) error {
	return nil
}
