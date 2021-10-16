package doctor

import (
	"errors"
	"log"

	"github.com/zakiafada32/vaccination-record/business"
)

type service struct {
	repository Repository
}

func NewDoctorService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FindAllDoctor() ([]Doctor, error) {
	users, err := s.repository.FindAllDoctor()
	if err != nil {
		log.Println(err)
		return []Doctor{}, errors.New(business.InternalServerError)
	}
	return users, nil
}
