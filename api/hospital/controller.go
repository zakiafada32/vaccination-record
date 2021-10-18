package hospital

import (
	"github.com/labstack/echo/v4"

	message "github.com/zakiafada32/vaccination-record/business"
	hospitalBusiness "github.com/zakiafada32/vaccination-record/business/hospital"
	"github.com/zakiafada32/vaccination-record/utils"
)

type HospitalController struct {
	service hospitalBusiness.Service
}

func NewHospitalController(service hospitalBusiness.Service) *HospitalController {
	return &HospitalController{
		service: service,
	}
}

func (dc *HospitalController) FindAll(c echo.Context) error {
	hospitals, err := dc.service.FindAllHospital()
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.Success, echo.Map{
		"hospitals": hospitals,
	}))
}
