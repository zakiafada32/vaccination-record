package modules

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	userBusiness "github.com/zakiafada32/vaccination-record/business/user"
	vaccineBusiness "github.com/zakiafada32/vaccination-record/business/vaccine"
	"github.com/zakiafada32/vaccination-record/utils"
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

func (repo *userRepository) RegisterUser(data userBusiness.User) error {
	userData := convertToUserModel(data)
	err := repo.db.Create(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindAllUser() ([]userBusiness.User, error) {
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

func (repo *userRepository) FindUserById(id string) (userBusiness.User, error) {
	var userData User
	err := repo.db.Where("id = ?", id).First(&userData).Error
	if err != nil {
		return userBusiness.User{}, err
	}
	user := convertToUserBusiness(userData)
	return user, nil
}

func (repo *userRepository) FindUserByIdentityCardNumber(cardId string) (userBusiness.User, error) {
	var userData User
	err := repo.db.Where("identity_card_number = ?", cardId).First(&userData).Error
	if err != nil {
		return userBusiness.User{}, err
	}
	user := convertToUserBusiness(userData)
	return user, nil
}

func (repo *userRepository) FindUserByEmail(email string) error {
	var userData User
	err := repo.db.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindUserByPhoneNumber(phoneNumber string) error {
	var userData User
	err := repo.db.Where("phone_number = ?", phoneNumber).First(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindLatestVaccineHistoryOfUser(userId string) (vaccineBusiness.Vaccine, error) {
	var vaccine Vaccine
	err := repo.db.Order("vaccinated_date desc").Preload(clause.Associations).Where("user_id = ?", userId).First(&vaccine).Error
	if err != nil {
		return vaccineBusiness.Vaccine{}, err
	}

	vaccineData := convertToVaccineBusiness(vaccine)
	return vaccineData, nil
}

func (repo *userRepository) FindVaccinesByUserId(userId string) ([]vaccineBusiness.Vaccine, error) {
	var vaccines []Vaccine
	err := repo.db.Order("vaccinated_date desc").Preload(clause.Associations).Where("user_id = ?", userId).Find(&vaccines).Error
	if err != nil {
		return []vaccineBusiness.Vaccine{}, err
	}

	vaccinesData := make([]vaccineBusiness.Vaccine, len(vaccines))
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
