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

type EmployeeFactory func(int) IEmployee

var employeeFactories = make(map[string]EmployeeFactory)

func RegisterEmployee(role string, factory EmployeeFactory) {
	employeeFactories[role] = factory
}

func GetEmployee(index int, role string) (IEmployee, error) {
	if factory, ok := employeeFactories[role]; ok {
		return factory(index), nil
	}
	return nil, helper.ErrInvalidEmployee
}
