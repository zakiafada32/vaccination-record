package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/zakiafada32/vaccination-record/api/doctor"
	"github.com/zakiafada32/vaccination-record/api/hospital"
	"github.com/zakiafada32/vaccination-record/api/user"
	"github.com/zakiafada32/vaccination-record/api/vaccine"
	"github.com/zakiafada32/vaccination-record/utils"
)

type Controller struct {
	User     *user.UserController
	Vaccine  *vaccine.VaccineController
	Doctor   *doctor.DoctorController
	Hospital *hospital.HospitalController
}

func Bootstrap(e *echo.Echo, c *Controller) {
	if c.User == nil {
		panic("user controller cannot be nil")
	}

	if c.Vaccine == nil {
		panic("vaccine controller cannot be nil")
	}

	if c.Doctor == nil {
		panic("doctor controller cannot be nil")
	}

	if c.Hospital == nil {
		panic("hospital controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	routeV1 := e.Group("api/vaccine-user")

	routeV1.POST("/register", c.User.Register)
	routeV1.GET("", c.User.FindAll)

	routeV1.POST("/:userId/add-history", c.Vaccine.AddHistory)
	routeV1.DELETE("/:userId/del-history/:vaccineId", c.Vaccine.DeleteHistory)
}
