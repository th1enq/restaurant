package food

const (
	PENDING = "pending"
	MAKING  = "making"
	SERVED  = "served"
)

const (
	PIZZA  = "pizza"
	PASTA  = "pasta"
	BURGER = "burger"
	NONE   = "none"
)

type Food struct {
	status string
}

type IFood interface {
	GetName() string
	GetStatus() string
	SetStatus(string) error
}
