package modules

import (
	"gorm.io/gorm"

	doctorBusiness "github.com/zakiafada32/vaccination-record/business/doctor"
)

type Doctor struct {
	ID        uint32 `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:255"`
	StrNumber string `gorm:"not null;size:20;unique"`
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *doctorRepository {
	return &doctorRepository{
		db: db,
	}
}

func (repo *doctorRepository) FindAllDoctor() ([]doctorBusiness.Doctor, error) {
	var doctorsData []Doctor
	err := repo.db.Find(&doctorsData).Error
	if err != nil {
		return []doctorBusiness.Doctor{}, err
	}

	doctors := make([]doctorBusiness.Doctor, len(doctorsData))
	for i, doctor := range doctorsData {
		doctors[i] = convertToDoctorBusiness(doctor)
	}

	return doctors, nil
}

func convertToDoctorModel(data doctorBusiness.Doctor) Doctor {
	return Doctor{
		Name:      data.Name,
		StrNumber: data.StrNumber,
	}
}

func convertToDoctorBusiness(data Doctor) doctorBusiness.Doctor {
	return doctorBusiness.Doctor{
		ID:        data.ID,
		Name:      data.Name,
		StrNumber: data.StrNumber,
	}
}
