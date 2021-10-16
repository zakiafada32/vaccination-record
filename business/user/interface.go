package user

type Service interface {
	FindAll() ([]User, error)
	Create(user User) (id string, err error)
	AddHistory(userId string, vaccine Vaccine) ([]Vaccine, error)
}

type Repository interface {
	FindById(userId string) (User, error)
	FindByIdentityCardNumber(cardId string) error
	FindByEmail(email string) error
	FindByPhoneNumber(phoneNumber string) error
	FindAll() ([]User, error)
	Create(user User) error
	FindOrAddHospital(hospital Hospital) (hospitalId uint32, err error)
	FindOrAddDoctor(doctor Doctor) (doctorId uint32, err error)
	AddHistoryVaccine(userId string, vaccine Vaccine) error
	FindVaccinesByUserId(userId string) ([]Vaccine, error)
}
