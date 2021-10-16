package modules

import (
	"gorm.io/gorm"

	hospitalBusiness "github.com/zakiafada32/vaccination-record/business/hospital"
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

func (repo *hospitalRepository) FindAllHospital() ([]hospitalBusiness.Hospital, error) {
	var hospitalsData []Hospital
	err := repo.db.Find(&hospitalsData).Error
	if err != nil {
		return []hospitalBusiness.Hospital{}, err
	}

	hospitals := make([]hospitalBusiness.Hospital, len(hospitalsData))
	for i, hospital := range hospitalsData {
		hospitals[i] = convertToHospitalBusiness(hospital)
	}

	return hospitals, nil
}

func convertToHospitalModel(data hospitalBusiness.Hospital) Hospital {
	return Hospital{
		Name:    data.Name,
		Address: data.Address,
	}
}

func convertToHospitalBusiness(data Hospital) hospitalBusiness.Hospital {
	return hospitalBusiness.Hospital{
		Name:    data.Name,
		Address: data.Address,
	}
}
