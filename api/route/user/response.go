package user

import userBusiness "github.com/zakiafada32/vaccine/business/user"

type userResponse struct {
	ID                 string `json:"id"`
	IdentityCardNumber string `json:"identity_card_number"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	Dob                string `json:"dob"`
	Address            string `json:"address"`
}

func convertToUserResponse(user userBusiness.User) userResponse {
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
