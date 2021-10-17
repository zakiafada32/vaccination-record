package user

import userBusiness "github.com/zakiafada32/vaccination-record/business/user"

type registerUserRequestBody struct {
	IdentityCardNumber string `json:"identity_card_number" validate:"required,gt=15"`
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"required,email"`
	PhoneNumber        string `json:"phone_number" validate:"required,gt=9"`
	Dob                string `json:"dob" validate:"required,datetime=2006-01-02"`
	Address            string `json:"address" validate:"required"`
	Password           string `json:"password" validate:"required,gt=6"`
}

func (req *registerUserRequestBody) convertToUserBusiness() userBusiness.User {
	return userBusiness.User{
		IdentityCardNumber: req.IdentityCardNumber,
		Name:               req.Name,
		Email:              req.Email,
		PhoneNumber:        req.PhoneNumber,
		Dob:                req.Dob,
		Address:            req.Address,
		Password:           req.Password,
	}
}

type checkStatusRequestBody struct {
	IdentityCardNumber string `json:"identity_card_number" validate:"required,gt=15"`
	Name               string `json:"name" validate:"required"`
	Password           string `json:"password" validate:"required,gt=6"`
}
