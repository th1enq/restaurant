package employee

import (
	"restaurant/helper"
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
}

type IEmployee interface {
	Work(chan<- interface{}, *sync.WaitGroup, string)
	SetStatus(string)
}

func (c *Employee) SetStatus(status string) {
	c.Status = status
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
