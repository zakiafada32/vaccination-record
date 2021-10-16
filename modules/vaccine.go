package modules

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	doctorBusiness "github.com/zakiafada32/vaccination-record/business/doctor"
	hospitalBusiness "github.com/zakiafada32/vaccination-record/business/hospital"
	vaccineBusiness "github.com/zakiafada32/vaccination-record/business/vaccine"
	"github.com/zakiafada32/vaccination-record/utils"
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

func (repo *vaccineRepository) FindOrAddDoctor(data doctorBusiness.Doctor) (doctorId uint32, err error) {
	doctorData := convertToDoctorModel(data)
	err = repo.db.FirstOrCreate(&doctorData, Doctor{StrNumber: doctorData.StrNumber}).Error
	if err != nil {
		return 0, err
	}
	return doctorData.ID, nil
}

func (repo *vaccineRepository) FindOrAddHospital(data hospitalBusiness.Hospital) (hospitalId uint32, err error) {
	hospitalData := convertToHospitalModel(data)
	err = repo.db.FirstOrCreate(&hospitalData, Hospital{Name: hospitalData.Name}).Error
	if err != nil {
		return 0, err
	}
	return hospitalData.ID, nil
}

func (repo *vaccineRepository) AddHistoryVaccine(userId string, data vaccineBusiness.Vaccine) error {
	vaccineData := convertToVaccineModel(userId, data)
	err := repo.db.Create(&vaccineData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *vaccineRepository) FindVaccinesByUserId(userId string) ([]vaccineBusiness.VaccineResponse, error) {
	var vaccines []Vaccine
	err := repo.db.Order("vaccinated_date desc").Preload(clause.Associations).Where("user_id = ?", userId).Find(&vaccines).Error
	if err != nil {
		return []vaccineBusiness.VaccineResponse{}, err
	}

	vaccinesData := make([]vaccineBusiness.VaccineResponse, len(vaccines))
	for i, vaccine := range vaccines {
		vaccinesData[i] = convertToVaccineBusiness(vaccine)
	}

	return vaccinesData, nil
}

func (repo *vaccineRepository) FindVaccinesByVaccineIdAndUserId(userId string, vaccineId uint32) error {
	var vaccine Vaccine
	err := repo.db.Where("id = ? AND user_id = ?", vaccineId, userId).First(&vaccine).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *vaccineRepository) DeleteVaccineByVaccineIdAndUserId(userId string, vaccineId uint32) error {
	err := repo.db.Where("user_id = ?", userId).Delete(&User{}, vaccineId).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *vaccineRepository) FindUserById(userId string) error {
	var userData User
	err := repo.db.Where("id = ?", userId).Find(&userData).Error
	if err != nil {
		return err
	}

	return nil
}

func convertToVaccineModel(userId string, data vaccineBusiness.Vaccine) Vaccine {
	return Vaccine{
		Description:    data.Description,
		VaccinatedDate: utils.ConvertStringToDate(data.VaccinatedDate),
		UserID:         userId,
		DoctorID:       data.Doctor.ID,
		HospitalID:     data.Hospital.ID,
	}
}

func convertToVaccineBusiness(data Vaccine) vaccineBusiness.VaccineResponse {
	return vaccineBusiness.VaccineResponse{
		ID:             data.ID,
		Description:    data.Description,
		VaccinatedDate: utils.ConvertDateToString(data.VaccinatedDate),
		Hospital: vaccineBusiness.HospitalResponse{
			Name:    data.Hospital.Name,
			Address: data.Hospital.Address,
		},
		Doctor: vaccineBusiness.DoctorResponse{
			Name:      data.Doctor.Name,
			StrNumber: data.Doctor.StrNumber,
		},
	}
}
