package main

import (
	"fmt"
	"restaurant/customer"
	"restaurant/drinking"
	"restaurant/employee"
	"restaurant/food"
	"restaurant/order"
	"sync"
	"time"
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
	wg            sync.WaitGroup
	readyFood     = make(chan interface{})
	readyDrinking = make(chan interface{})
)

var orderHistory sync.Map

func main() {
	// Initialize chefs, bartenders, and waiters
	initializeEmployees()

	// Customers order food and drinks
	placeOrder()

	// Fan-in pattern: Collect results from readyFood and readyDrinking into announcement
	announcement := fanIn(readyFood, readyDrinking)

	var foodOrderList = make(chan order.Order, len(manager.FoodLists))
	var drinkOrderList = make(chan order.Order, len(manager.DrinkLists))

	for _, chef := range chefs {
		wg.Add(1)
		go chef.Work(readyFood, &wg, foodOrderList)
	}

	for _, bartender := range bartenders {
		wg.Add(1)
		go bartender.Work(readyDrinking, &wg, drinkOrderList)
	}

	for _, order := range manager.FoodLists {
		fmt.Println("Order: ", order)
		foodOrderList <- order
	}
	for _, order := range manager.DrinkLists {
		fmt.Println("Order: ", order)
		foodOrderList <- order
	}
	close(foodOrderList)
	close(drinkOrderList)

	// Wait for chefs and bartenders to finish
	go func() {
		wg.Wait()
		close(readyFood)
		close(readyDrinking)
	}()

	// Fan-out pattern: Waiters serve the food and drinks
	waitForWaiters(announcement)

	fmt.Println("All waiters finished serving!")

	// orderTrace()
}

func initializeEmployees() {
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
}

func placeOrder() {
	Luffy.Order(food.PIZZA)
	Luffy.Order(food.PASTA)
	Luffy.Order(drinking.TEA)

	Zoro.Order(drinking.COFFEE)
	Zoro.Order(food.PASTA)
	Zoro.Order(food.BURGER)

	Chopper.Order(food.PIZZA)
	Chopper.Order(food.BURGER)
	Chopper.Order(drinking.JUICE)

	Nami.Order(food.BURGER)
	Nami.Order(drinking.JUICE)
	Nami.Order(food.PASTA)
}

func fanIn(readyFood, readyDrinking <-chan interface{}) <-chan interface{} {
	announcement := make(chan interface{})
	go func() {
		defer close(announcement)
		for val := range manager.Listen(readyFood, readyDrinking) {
			announcement <- val
			fmt.Println(val)
			orderHistory.Store(val, time.Now())
		}
	}()
	return announcement
}

func waitForWaiters(announcement <-chan interface{}) {
	var wg2 sync.WaitGroup
	for _, waiter := range waiters {
		wg2.Add(1)
		go func(w *employee.Waiter) {
			w.Work(announcement, &wg2)
		}(waiter)
	}
	wg2.Wait()
}

// func orderTrace() {
// 	fmt.Println("Restaurant history:")
// 	orderHistory.Range(func(key, value interface{}) bool {
// 		order := key.(*order.Order)

// 		timestamp := value.(time.Time).Format("2006-01-02 15:04:05.000000000")

// 		if foodItem, ok := order.Item.(food.IFood); ok {
// 			fmt.Printf("Customer: %s ordered %s at %v\n", order.NameCustomer, foodItem.GetFoodName(), timestamp)
// 		} else {
// 			drinkItem := order.Item.(drinking.IDrinking)
// 			fmt.Printf("Customer: %s ordered %s at %v\n", order.NameCustomer, drinkItem.GetDrinkingName(), timestamp)
// 		}
// 		return true
// 	})
// }
