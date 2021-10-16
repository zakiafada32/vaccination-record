package vaccine

import (
	"strconv"

	"github.com/labstack/echo/v4"

	message "github.com/zakiafada32/vaccination-record/business"
	vaccineBusiness "github.com/zakiafada32/vaccination-record/business/vaccine"
	"github.com/zakiafada32/vaccination-record/utils"
)

type VaccineController struct {
	service vaccineBusiness.Service
}

func NewVaccineController(service vaccineBusiness.Service) *VaccineController {
	return &VaccineController{
		service: service,
	}
}

func (uc *VaccineController) AddHistory(c echo.Context) error {
	userId := c.Param("userId")
	var body addHistoryRequestBody
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(utils.ConstructResponse(message.BadRequest, echo.Map{}))
	}

	err = c.Validate(&body)
	if err != nil {
		return err
	}

	vaccines, err := uc.service.AddHistory(userId, body.convertToVaccineBusiness())
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.SuccessCreated, echo.Map{
		"vaccines": vaccines,
	}))
}

func (uc *VaccineController) DeleteHistory(c echo.Context) error {
	userId := c.Param("userId")
	vaccineIdStr := c.Param("vaccineId")

	vaccineId, err := strconv.ParseUint(vaccineIdStr, 10, 32)
	if err != nil {
		return c.JSON(utils.ConstructResponse(message.NotFound, echo.Map{}))
	}

	vaccines, err := uc.service.DeleteHistory(userId, uint32(vaccineId))
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.SuccessCreated, echo.Map{
		"vaccines": vaccines,
	}))
}
