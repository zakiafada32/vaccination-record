package doctor

type Service interface {
	FindAllDoctor() ([]Doctor, error)
}

type Repository interface {
	FindAllDoctor() ([]Doctor, error)
}
