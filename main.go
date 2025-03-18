package main

import (
	"fmt"
	"restaurant/customer"
	"restaurant/drinking"
	"restaurant/employee"
	"restaurant/food"
	"sync"
)

const (
	NUMBEROFCHEFS      = 3
	NUMBEROFBARTENDERS = 3
	NUMBEROFWAITERS    = 3
)

var (
	Luffy   = customer.NewCustomer("Luffy")
	Zoro    = customer.NewCustomer("Zoro")
	Chopper = customer.NewCustomer("Chopper")
	Nami    = customer.NewCustomer("Nami")
)

var (
	manager    = employee.GetManager()
	chefs      = make([]*employee.Chef, NUMBEROFCHEFS)
	bartenders = make([]*employee.Bartender, NUMBEROFBARTENDERS)
	waiters    = make([]*employee.Waiter, NUMBEROFWAITERS)
)

var (
	readyFood     = make(chan interface{})
	readyDrinking = make(chan interface{})
)

func main() {
	// Initialize chefs, bartenders, and waiters
	for i := 0; i < NUMBEROFCHEFS; i++ {
		e, _ := employee.GetEmployee(i+1, employee.CHEF)
		chefs[i] = e.(*employee.Chef)
	}
	for i := 0; i < NUMBEROFBARTENDERS; i++ {
		e, _ := employee.GetEmployee(i+1, employee.BARTENDER)
		bartenders[i] = e.(*employee.Bartender)
	}
	for i := 0; i < NUMBEROFWAITERS; i++ {
		waiters[i] = employee.NewWaiter(i + 1)
	}

	// Customers order food and drinks
	Luffy.Order(food.PIZZA)
	// Luffy.Order(drinking.COFFEE)

	Zoro.Order(food.PASTA)
	// Zoro.Order(drinking.TEA)

	Chopper.Order(food.PIZZA)
	// Chopper.Order(drinking.JUICE)

	Nami.Order(food.PASTA)
	// Nami.Order(drinking.JUICE)

	var wg sync.WaitGroup
	orderLists := manager.OrderLists

	// Process orders
	for _, order := range orderLists {
		wg.Add(1)
		assigned := false

		for !assigned {
			foodOrder, err := food.GetFood(order.ThingName)
			if err == nil {
				for _, chef := range chefs {
					chef.Mu.Lock()
					if chef.Status == employee.RELAX {
						chef.Status = "Working"
						chef.Mu.Unlock()
						go chef.Work(readyFood, &wg, foodOrder.GetFoodName())
						assigned = true
						break
					}
					chef.Mu.Unlock()
				}
			} else {
				drinkingOrder, _ := drinking.GetDrinking(order.ThingName)
				for _, bartender := range bartenders {
					bartender.Mu.Lock()
					if bartender.Status == employee.RELAX {
						go bartender.Work(readyDrinking, &wg, drinkingOrder.GetDrinkingName())
						assigned = true
						break
					}
					bartender.Mu.Unlock()
				}
			}
		}
	}

	// Wait for chefs and bartenders to finish
	go func() {
		wg.Wait()
		close(readyFood)
		close(readyDrinking)
	}()

	announcement := make(chan interface{})

	// Fan-in pattern: Collect results from readyFood and readyDrinking into announcement
	go func() {
		for val := range manager.Listen(readyFood, readyDrinking) {
			announcement <- val
		}
		close(announcement)
	}()

	// Fan-out pattern: Waiters serve the food and drinks
	var wg2 sync.WaitGroup
	for _, waiter := range waiters {
		wg2.Add(1)
		go func(w *employee.Waiter) {
			w.Work(announcement, &wg2)
		}(waiter)
	}
	wg2.Wait()

	fmt.Println("All waiters finished serving!")
}
