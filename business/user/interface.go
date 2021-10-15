package user

type Service interface {
	FindAll() (users []User, err error)
	Create(user User) (id string, err error)
}

type Repository interface {
	FindById(id string) (user User, err error)
	FindAll() (users []User, err error)
	Create(user User) error
}
