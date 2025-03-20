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

const READY = true
const BUSY = false

type Employee struct {
	Ready chan interface{}
	ID    int
}

type IEmployee interface {
	Work(chan<- interface{}, *sync.WaitGroup, chan order.Order, *sync.Map)
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
