package modules

import (
	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	"gorm.io/gorm"
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

func convertToDoctorModel(data userBusiness.Doctor) Doctor {
	return Doctor{
		Name:      data.Name,
		StrNumber: data.StrNumber,
	}
}
