package user

type Service interface {
	FindAll() ([]User, error)
	Create(user User) (id string, err error)
	// AddHistory(vaccine Vaccine) ([]Vaccine, error)
}

type Repository interface {
	FindById(id string) (User, error)
	FindByIdentityCardNumber(idCard string) error
	FindByEmail(email string) error
	FindByPhoneNumber(phoneNumber string) error
	FindAll() ([]User, error)
	Create(user User) error
	// FindHospitalByName(name string) (Hospital, error)
	// FindDoctorByStrNumber(strNumber string) (Doctor, error)
}
