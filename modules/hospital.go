package modules

import (
	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	"gorm.io/gorm"
)

type Hospital struct {
	ID      uint32 `gorm:"primaryKey"`
	Name    string `gorm:"not null;size:255;unique"`
	Address string `gorm:"not null;size:255"`
}

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) *hospitalRepository {
	return &hospitalRepository{
		db: db,
	}
}

func convertToHospitalModel(data userBusiness.Hospital) Hospital {
	return Hospital{
		Name:    data.Name,
		Address: data.Address,
	}
}
