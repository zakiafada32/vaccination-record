package doctor

import (
	"github.com/labstack/echo/v4"

	message "github.com/zakiafada32/vaccination-record/business"
	doctorBusiness "github.com/zakiafada32/vaccination-record/business/doctor"
	"github.com/zakiafada32/vaccination-record/utils"
)

type DoctorController struct {
	service doctorBusiness.Service
}

func NewDoctorController(service doctorBusiness.Service) *DoctorController {
	return &DoctorController{
		service: service,
	}
}

func (dc *DoctorController) FindAll(c echo.Context) error {
	doctors, err := dc.service.FindAllDoctor()
	if err != nil {
		return c.JSON(utils.ConstructResponse(err.Error(), echo.Map{}))
	}

	return c.JSON(utils.ConstructResponse(message.Success, echo.Map{
		"doctors": doctors,
	}))
}
