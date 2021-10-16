package hospital

type Service interface {
	FindAllHospital() ([]Hospital, error)
}

type Repository interface {
	FindAllHospital() ([]Hospital, error)
}
