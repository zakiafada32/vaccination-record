package user

import "github.com/zakiafada32/vaccination-record/business/vaccine"

type User struct {
	ID                 string `json:"user_unique_id"`
	IdentityCardNumber string `json:"identity_card_number" validate:"required,gt=15"`
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"required,email"`
	PhoneNumber        string `json:"phone_number" validate:"required,gt=9"`
	Dob                string `json:"dob" validate:"required,datetime=2006-01-02"`
	Address            string `json:"address" validate:"required"`
	Password           string `json:"password" validate:"required,gt=6"`
}

type userResponse struct {
	ID                 string `json:"user_unique_id"`
	IdentityCardNumber string `json:"identity_card_number"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	Dob                string `json:"dob"`
	Address            string `json:"address"`
}

type checkStatus struct {
	ID                 string                  `json:"user_unique_id"`
	IdentityCardNumber string                  `json:"identity_card_number"`
	Name               string                  `json:"name"`
	Vaccine            vaccine.VaccineResponse `json:"vaccine"`
}
