package user

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
	"github.com/zakiafada32/vaccination-record/utils"
)

type service struct {
	repository Repository
}

func NewUserService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) RegisterUser(user User) (id string, err error) {
	err = utils.GetValidatorStruct().Struct(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.BadRequest)
	}

	err = s.repository.FindUserByIdentityCardNumber(user.IdentityCardNumber)
	if err == nil {
		log.Println("identity card number already existed")
		return "", errors.New(business.IdentityCardNumberAlreadyExist)
	}

	err = s.repository.FindUserByPhoneNumber(user.PhoneNumber)
	if err == nil {
		log.Println("phone number already existed")
		return "", errors.New(business.PhoneNumberAlreadyExist)
	}

	err = s.repository.FindUserByEmail(user.Email)
	if err == nil {
		log.Println("email alredy exist")
		return "", errors.New(business.EmailAlreadyExist)
	}

	user.ID = utils.GenerateID()
	user.Password, err = utils.Hashing(user.Password)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}

	err = s.repository.RegisterUser(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}
	return user.ID, nil
}

func (s *service) FindAllUser() ([]User, error) {
	users, err := s.repository.FindAllUser()
	if err != nil {
		log.Println(err)
		return []User{}, errors.New(business.InternalServerError)
	}
	return users, nil
}

func (s *service) FindById(id string) (User, error) {
	user, err := s.repository.FindUserById(id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.NotFound)
	}
	return user, nil
}
