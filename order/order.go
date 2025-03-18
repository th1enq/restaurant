package order

type Order struct {
	ThingName    string
	NameCustomer string
}

func NewOrder(thing string, name string) *Order {
	return &Order{
		ThingName:    thing,
		NameCustomer: name,
	}
}
