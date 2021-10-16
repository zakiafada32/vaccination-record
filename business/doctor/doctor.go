package doctor

type Doctor struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name" validate:"required"`
	StrNumber string `json:"str_number" validate:"required"`
}
