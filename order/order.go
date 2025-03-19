package order

type Order struct {
	Things       interface{}
	NameCustomer string
}

func NewOrder(Thing interface{}, name string) *Order {
	return &Order{
		Things:       Thing,
		NameCustomer: name,
	}
}
