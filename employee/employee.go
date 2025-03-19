package employee

import (
	"restaurant/helper"
	"restaurant/order"
	"sync"
)

const (
	CHEF      = "chef"
	BARTENDER = "bartender"
	WAITER    = "waiter"
)

const RELAX = "relax"

type Employee struct {
	Status string
	ID     int
	Mu     sync.Mutex
}

type IEmployee interface {
	Work(chan<- interface{}, *sync.WaitGroup, order.Order)
}

func GetEmployee(index int, role string) (IEmployee, error) {
	switch role {
	case CHEF:
		return newChef(index), nil
	case BARTENDER:
		return newBartender(index), nil
	default:
		return nil, helper.ErrInvalidEmployee
	}
}
