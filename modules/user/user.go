package user

import (
	"time"

	userBusiness "github.com/zakiafada32/vaccine/business/user"
	"github.com/zakiafada32/vaccine/modules/utils"
	"gorm.io/gorm"
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

func (repo *userRepository) Create(user userBusiness.User) error {
	userData := convertToUserModel(user)
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

func convertToUserModel(user userBusiness.User) User {
	return User{
		ID:                 user.ID,
		IdentityCardNumber: user.IdentityCardNumber,
		Name:               user.Name,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Dob:                utils.ConvertStringToDate(user.Dob),
		Password:           user.Password,
		Address:            user.Address,
	}
}

func convertToUserBusiness(user User) userBusiness.User {
	return userBusiness.User{
		ID:                 user.ID,
		IdentityCardNumber: user.IdentityCardNumber,
		Name:               user.Name,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Dob:                utils.ConvertDateToString(user.Dob),
		Password:           user.Password,
		Address:            user.Address,
	}
}
