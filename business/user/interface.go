package user

import "github.com/zakiafada32/vaccination-record/business/vaccine"

type Service interface {
	FindAllUser() ([]userResponse, error)
	RegisterUser(user User) (id string, err error)
	CheckStatus(idCard, name, password string) (checkStatus, error)
}

type Repository interface {
	FindUserById(userId string) (User, error)
	FindUserByIdentityCardNumber(cardId string) (User, error)
	FindUserByEmail(email string) error
	FindUserByPhoneNumber(phoneNumber string) error
	FindAllUser() ([]User, error)
	RegisterUser(user User) error
	FindLatestVaccineHistoryOfUser(userId string) (vaccine.Vaccine, error)
}
