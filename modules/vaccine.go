package modules

import (
	"time"

	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	"github.com/zakiafada32/vaccination-record/utils"

	"gorm.io/gorm"
)

type Vaccine struct {
	ID             uint32    `gorm:"primaryKey"`
	Description    string    `gorm:"not null;size:255"`
	VaccinatedDate time.Time `gorm:"not null"`
	UserID         string    `gorm:"not null;size:36"`
	User           User
	HospitalID     uint32 `gorm:"not null"`
	Hospital       Hospital
	DoctorID       uint32 `gorm:"not null"`
	Doctor         Doctor
}

type vaccineRepository struct {
	db *gorm.DB
}

func NewVaccineRepository(db *gorm.DB) *vaccineRepository {
	return &vaccineRepository{
		db: db,
	}
}

func convertToVaccineModel(userId string, data userBusiness.Vaccine) Vaccine {
	return Vaccine{
		Description:    data.Description,
		VaccinatedDate: utils.ConvertStringToDate(data.VaccinatedDate),
		UserID:         userId,
		DoctorID:       data.Doctor.ID,
		HospitalID:     data.Hospital.ID,
	}
}

func convertToVaccineBusiness(data Vaccine) userBusiness.Vaccine {
	return userBusiness.Vaccine{
		ID:             data.ID,
		Description:    data.Description,
		VaccinatedDate: utils.ConvertDateToString(data.VaccinatedDate),
		Hospital: userBusiness.Hospital{
			ID:      data.Hospital.ID,
			Name:    data.Hospital.Name,
			Address: data.Hospital.Address,
		},
		Doctor: userBusiness.Doctor{
			ID:        data.Doctor.ID,
			Name:      data.Doctor.Name,
			StrNumber: data.Doctor.StrNumber,
		},
	}
}
