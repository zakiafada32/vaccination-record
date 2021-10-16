package modules

import (
	"time"

	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	"github.com/zakiafada32/vaccination-record/modules/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID                 string    `gorm:"primaryKey;size:36"`
	IdentityCardNumber string    `gorm:"not null;unique;size:16"`
	Name               string    `gorm:"not null;size:255"`
	Email              string    `gorm:"not null;unique;size:255"`
	PhoneNumber        string    `gorm:"not null;unique;size:16"`
	Dob                time.Time `gorm:"not null"`
	Address            string    `gorm:"not null;size:255"`
	Password           string    `gorm:"not null;size:255"`
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) Create(data userBusiness.User) error {
	userData := convertToUserModel(data)
	err := repo.db.Create(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindAll() ([]userBusiness.User, error) {
	var usersData []User
	err := repo.db.Find(&usersData).Error
	if err != nil {
		return []userBusiness.User{}, err
	}

	users := make([]userBusiness.User, len(usersData))
	for i, user := range usersData {
		users[i] = convertToUserBusiness(user)
	}

	return users, nil
}

func (repo *userRepository) FindById(id string) (userBusiness.User, error) {
	var userData User
	err := repo.db.Where("id = ?", id).First(&userData).Error
	if err != nil {
		return userBusiness.User{}, err
	}
	user := convertToUserBusiness(userData)
	return user, nil
}

func (repo *userRepository) FindByIdentityCardNumber(idCard string) error {
	var userData User
	err := repo.db.Where("identity_card_number = ?", idCard).First(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindByEmail(email string) error {
	var userData User
	err := repo.db.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindByPhoneNumber(phoneNumber string) error {
	var userData User
	err := repo.db.Where("phone_number = ?", phoneNumber).First(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindOrAddDoctor(data userBusiness.Doctor) (doctorId uint32, err error) {
	doctorData := convertToDoctorModel(data)
	err = repo.db.FirstOrCreate(&doctorData, Doctor{StrNumber: doctorData.StrNumber}).Error
	if err != nil {
		return 0, err
	}
	return doctorData.ID, nil
}

func (repo *userRepository) FindOrAddHospital(data userBusiness.Hospital) (hospitalId uint32, err error) {
	hospitalData := convertToHospitalModel(data)
	err = repo.db.FirstOrCreate(&hospitalData, Hospital{Name: hospitalData.Name}).Error
	if err != nil {
		return 0, err
	}
	return hospitalData.ID, nil
}

func (repo *userRepository) AddHistoryVaccine(userId string, data userBusiness.Vaccine) error {
	vaccineData := convertToVaccineModel(userId, data)
	err := repo.db.Create(&vaccineData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindVaccinesByUserId(userId string) ([]userBusiness.Vaccine, error) {
	var vaccines []Vaccine
	err := repo.db.Order("vaccinated_date desc").Preload(clause.Associations).Where("user_id = ?", userId).Find(&vaccines).Error
	if err != nil {
		return []userBusiness.Vaccine{}, err
	}

	vaccinesData := make([]userBusiness.Vaccine, len(vaccines))
	for i, vaccine := range vaccines {
		vaccinesData[i] = convertToVaccineBusiness(vaccine)
	}

	return vaccinesData, nil
}

func convertToUserModel(data userBusiness.User) User {
	return User{
		ID:                 data.ID,
		IdentityCardNumber: data.IdentityCardNumber,
		Name:               data.Name,
		Email:              data.Email,
		PhoneNumber:        data.PhoneNumber,
		Dob:                utils.ConvertStringToDate(data.Dob),
		Password:           data.Password,
		Address:            data.Address,
	}
}

func convertToUserBusiness(data User) userBusiness.User {
	return userBusiness.User{
		ID:                 data.ID,
		IdentityCardNumber: data.IdentityCardNumber,
		Name:               data.Name,
		Email:              data.Email,
		PhoneNumber:        data.PhoneNumber,
		Dob:                utils.ConvertDateToString(data.Dob),
		Password:           data.Password,
		Address:            data.Address,
	}
}
