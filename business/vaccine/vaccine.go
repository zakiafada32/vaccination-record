package vaccine

import (
	"github.com/zakiafada32/vaccination-record/business/doctor"
	"github.com/zakiafada32/vaccination-record/business/hospital"
)

type Vaccine struct {
	ID             uint32            `json:"history_unique_id"`
	Description    string            `json:"description" validate:"required"`
	VaccinatedDate string            `json:"vaccinated_date" validate:"required,datetime=2006-01-02"`
	Hospital       hospital.Hospital `json:"hospital"`
	Doctor         doctor.Doctor     `json:"doctor"`
}

type VaccineResponse struct {
	ID             uint32           `json:"history_unique_id"`
	Description    string           `json:"description"`
	VaccinatedDate string           `json:"vaccinated_date"`
	Hospital       hospitalResponse `json:"hospital"`
	Doctor         doctorResponse   `json:"doctor"`
}

type hospitalResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type doctorResponse struct {
	Name      string `json:"name"`
	StrNumber string `json:"str_number"`
}
