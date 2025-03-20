package employee

import (
	"restaurant/order"
	"sync"
)

const NUMBEROFTHINGS = 2

type Manager struct {
	FoodLists  []order.Order
	DrinkLists []order.Order
}

var (
	instance *Manager
	once     sync.Once
)

func GetManager() *Manager {
	once.Do(func() {
		instance = &Manager{}
	})

	return instance
}

func (m *Manager) Listen(readyFood <-chan interface{}, readyDrinking <-chan interface{}) <-chan interface{} {
	announcement := make(chan interface{})
	go func() {
		defer close(announcement)
		for readyFood != nil || readyDrinking != nil {
			select {
			case food, ok := <-readyFood:
				if !ok {
					readyFood = nil
				} else {
					announcement <- food
				}
			case drink, ok := <-readyDrinking:
				if !ok {
					readyDrinking = nil
				} else {
					announcement <- drink
				}
			}
		}
	}()
	return announcement
}

func (m *Manager) AddFoodOrder(order order.Order) {
	m.FoodLists = append(m.FoodLists, order)
}

func (m *Manager) AddDrinkOrder(order order.Order) {
	m.DrinkLists = append(m.DrinkLists, order)
}
