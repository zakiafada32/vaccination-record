package user

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
	"github.com/zakiafada32/vaccination-record/business/vaccine"
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

	_, err = s.repository.FindUserByIdentityCardNumber(user.IdentityCardNumber)
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

func (s *service) FindAllUser() ([]userResponse, error) {
	users, err := s.repository.FindAllUser()
	if err != nil {
		log.Println(err)
		return []userResponse{}, errors.New(business.InternalServerError)
	}

	usersRespose := make([]userResponse, len(users))
	for i, user := range users {
		usersRespose[i] = convertToUserResponse(user)
	}
	return usersRespose, nil
}

func (s *service) FindById(id string) (User, error) {
	user, err := s.repository.FindUserById(id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New(business.NotFound)
	}
	return user, nil
}

func (s *service) CheckStatus(cardId, name, password string) (checkStatus, error) {
	user, err := s.repository.FindUserByIdentityCardNumber(cardId)
	if err != nil {
		log.Println(err)
		return checkStatus{}, errors.New(business.NotFound)
	}

	if user.Name != name {
		log.Println(err)
		return checkStatus{}, errors.New(business.NotFound)
	}

	err = utils.CompareHash(user.Password, password)
	if err != nil {
		log.Println(err)
		return checkStatus{}, errors.New(business.NotFound)
	}

	vaccineData, err := s.repository.FindVaccinesByUserId(user.ID)
	if err != nil {
		log.Println(err)
		return checkStatus{}, errors.New(business.InternalServerError)
	}

	var userStatus checkStatus
	if len(vaccineData) == 0 {
		userStatus = checkStatus{
			ID:                 user.ID,
			IdentityCardNumber: user.IdentityCardNumber,
			Name:               user.Name,
		}
	} else {
		userStatus = checkStatus{
			ID:                 user.ID,
			IdentityCardNumber: user.IdentityCardNumber,
			Name:               user.Name,
			Vaccine:            vaccine.ConvertToVaccineResponse(vaccineData[0]),
		}
	}

	return userStatus, nil
}

func convertToUserResponse(user User) userResponse {
	return userResponse{
		ID:                 user.ID,
		IdentityCardNumber: user.IdentityCardNumber,
		Name:               user.Name,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Dob:                user.Dob,
		Address:            user.Address,
	}
}
