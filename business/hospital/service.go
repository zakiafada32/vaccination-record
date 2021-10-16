package hospital

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
)

type service struct {
	repository Repository
}

func NewHospitalService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindAllHospital() ([]Hospital, error) {
	users, err := s.repository.FindAllHospital()
	if err != nil {
		log.Println(err)
		return []Hospital{}, errors.New(business.InternalServerError)
	}
	return users, nil
}
