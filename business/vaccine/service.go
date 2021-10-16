package vaccine

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
	"github.com/zakiafada32/vaccination-record/utils"
)

type service struct {
	repository Repository
}

func NewVaccineService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) AddHistory(userId string, vaccine Vaccine) ([]VaccineResponse, error) {
	err := utils.GetValidatorStruct().Struct(vaccine)
	if err != nil {
		log.Println(err)
		return []VaccineResponse{}, errors.New(business.BadRequest)
	}

	err = s.repository.FindUserById(userId)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.NotFound)
	}

	doctorId, err := s.repository.FindOrAddDoctor(vaccine.Doctor)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}
	vaccine.Doctor.ID = doctorId

	hospitalId, err := s.repository.FindOrAddHospital(vaccine.Hospital)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}
	vaccine.Hospital.ID = hospitalId

	err = s.repository.AddHistoryVaccine(userId, vaccine)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}

	vaccines, err := s.repository.FindVaccinesByUserId(userId)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}

	return vaccines, nil
}

func (s *service) DeleteHistory(userId string, vaccineId uint32) ([]VaccineResponse, error) {
	err := s.repository.FindVaccinesByVaccineIdAndUserId(userId, vaccineId)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.NotFound)
	}

	err = s.repository.DeleteVaccineByVaccineIdAndUserId(userId, vaccineId)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}

	vaccines, err := s.repository.FindVaccinesByUserId(userId)
	if err != nil {
		return []VaccineResponse{}, errors.New(business.InternalServerError)
	}

	return vaccines, nil

}
