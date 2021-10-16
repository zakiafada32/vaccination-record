package vaccine

import (
	doctorBusiness "github.com/zakiafada32/vaccination-record/business/doctor"
	hospitalBusiness "github.com/zakiafada32/vaccination-record/business/hospital"
	vaccineBusiness "github.com/zakiafada32/vaccination-record/business/vaccine"
)

type addHistoryRequestBody struct {
	Description    string `json:"description"`
	VaccinatedDate string `json:"vaccinated_date"`
	Hospital       struct {
		Name    string `json:"name" validate:"required"`
		Address string `json:"address" validate:"required"`
	}
	Doctor struct {
		Name      string `json:"name" validate:"required"`
		StrNumber string `json:"str_number" validate:"required"`
	}
}

func (req *addHistoryRequestBody) convertToVaccineBusiness() vaccineBusiness.Vaccine {
	return vaccineBusiness.Vaccine{
		Description:    req.Description,
		VaccinatedDate: req.VaccinatedDate,
		Hospital: hospitalBusiness.Hospital{
			Name:    req.Hospital.Name,
			Address: req.Hospital.Address,
		},
		Doctor: doctorBusiness.Doctor{
			Name:      req.Doctor.Name,
			StrNumber: req.Doctor.StrNumber,
		},
	}
}
