package user

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccine/business"
	"github.com/zakiafada32/vaccine/business/utils"
)

type service struct {
	repository Repository
}

func NewUserService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(user User) (id string, err error) {
	err = utils.GetValidator().Struct(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.BadRequest)
	}
	user.ID = utils.GenerateID()
	user.Password, err = utils.Hashing(user.Password)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}
	err = s.repository.Create(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}
	return user.ID, nil
}

func (s *service) FindAll() (users []User, err error) {
	users, err = s.repository.FindAll()
	if err != nil {
		log.Println(err)
		return []User{}, errors.New(business.InternalServerError)
	}
	return users, nil
}

func (s *service) FindById(id string) (user User, err error) {
	user, err = s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.NotFound)
	}
	return user, nil
}
