package employee

import "restaurant/helper"

const (
	CHEF      = "chef"
	BARTENDER = "bartender"
	MANAGER   = "manager"
)

const PENDING = "pending"

type Employee struct {
	Status string
}

type IEmployee interface {
	Work()
	SetStatus(string)
	GetReadyThing() chan interface{}
}

func GetEmployee(role string) (IEmployee, error) {
	switch role {
	case CHEF:
		return newChef(), nil
	case BARTENDER:
		return newBartender(), nil
	default:
		return nil, helper.ErrInvalidEmployee
	}
}
