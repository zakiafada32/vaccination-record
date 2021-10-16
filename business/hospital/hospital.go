package hospital

type Hospital struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
