package user

type Service interface {
	FindAllUser() ([]User, error)
	RegisterUser(user User) (id string, err error)
}

type Repository interface {
	FindUserById(userId string) (User, error)
	FindUserByIdentityCardNumber(cardId string) error
	FindUserByEmail(email string) error
	FindUserByPhoneNumber(phoneNumber string) error
	FindAllUser() ([]User, error)
	RegisterUser(user User) error
}
