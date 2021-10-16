package vaccine

import (
	"github.com/zakiafada32/vaccination-record/business/doctor"
	"github.com/zakiafada32/vaccination-record/business/hospital"
)

type Vaccine struct {
	ID             uint32 `json:"id"`
	Description    string `json:"description" validate:"required"`
	VaccinatedDate string `json:"vaccinated_date" validate:"required"`
	Hospital       hospital.Hospital
	Doctor         doctor.Doctor
}

type VaccineResponse struct {
	ID             uint32 `json:"id"`
	Description    string `json:"description"`
	VaccinatedDate string `json:"vaccinated_date"`
	Hospital       HospitalResponse
	Doctor         DoctorResponse
}

type HospitalResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DoctorResponse struct {
	Name      string `json:"name"`
	StrNumber string `json:"str_number"`
}
