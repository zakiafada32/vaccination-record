package user

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
	"github.com/zakiafada32/vaccination-record/business/utils"
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

	err = s.repository.FindByIdentityCardNumber(user.IdentityCardNumber)
	if err == nil {
		log.Println("identity card number already existed")
		return "", errors.New(business.IdentityCardNumberAlreadyExist)
	}

	err = s.repository.FindByPhoneNumber(user.PhoneNumber)
	if err == nil {
		log.Println("phone number already existed")
		return "", errors.New(business.PhoneNumberAlreadyExist)
	}

	err = s.repository.FindByEmail(user.Email)
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

	err = s.repository.Create(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(business.InternalServerError)
	}
	return user.ID, nil
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		log.Println(err)
		return []User{}, errors.New(business.InternalServerError)
	}
	return users, nil
}

func (s *service) FindById(id string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.NotFound)
	}
	return user, nil
}

func (s *service) AddHistory(userId string, vaccine Vaccine) ([]Vaccine, error) {
	err := utils.GetValidator().Struct(vaccine)
	if err != nil {
		log.Println(err)
		return []Vaccine{}, errors.New(business.BadRequest)
	}

	doctorId, err := s.repository.FindOrAddDoctor(vaccine.Doctor)
	if err != nil {
		return []Vaccine{}, errors.New(business.InternalServerError)
	}
	vaccine.Doctor.ID = doctorId

	hospitalId, err := s.repository.FindOrAddHospital(vaccine.Hospital)
	if err != nil {
		return []Vaccine{}, errors.New(business.InternalServerError)
	}
	vaccine.Hospital.ID = hospitalId

	err = s.repository.AddHistoryVaccine(userId, vaccine)
	if err != nil {
		return []Vaccine{}, errors.New(business.InternalServerError)
	}

	vaccines, err := s.repository.FindVaccinesByUserId(userId)
	if err != nil {
		return []Vaccine{}, errors.New(business.InternalServerError)
	}

	return vaccines, nil
}
