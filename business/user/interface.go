package user

type Service interface {
	FindAll() ([]User, error)
	Create(user User) (id string, err error)
}

type Repository interface {
	FindById(id string) (User, error)
	FindAll() ([]User, error)
	Create(user User) error
}
