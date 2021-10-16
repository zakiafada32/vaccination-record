package vaccine

import (
	"github.com/zakiafada32/vaccination-record/business/doctor"
	"github.com/zakiafada32/vaccination-record/business/hospital"
)

type Service interface {
	AddHistory(userId string, vaccine Vaccine) ([]VaccineResponse, error)
	DeleteHistory(userId string, vaccineId uint32) ([]VaccineResponse, error)
}

type Repository interface {
	FindUserById(userId string) error
	FindOrAddHospital(hospital hospital.Hospital) (hospitalId uint32, err error)
	FindOrAddDoctor(doctor doctor.Doctor) (doctorId uint32, err error)
	AddHistoryVaccine(userId string, vaccine Vaccine) error
	FindVaccinesByUserId(userId string) ([]VaccineResponse, error)
	FindVaccinesByVaccineIdAndUserId(userId string, vaccineId uint32) error
	DeleteVaccineByVaccineIdAndUserId(userId string, vaccineId uint32) error
}
