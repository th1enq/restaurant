package main

import (
	"fmt"
	"log"
	"restaurant/customer"
	"restaurant/drinking"
	"restaurant/employee"
	"restaurant/food"
	"restaurant/order"
	"sync"
)

const (
	NUMBEROFCHEFS      = 3
	NUMBEROFBARTENDERS = 3
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
)

var (
	wgFood        sync.WaitGroup
	wgDrinking    sync.WaitGroup
	readyFood     = make(chan interface{})
	readyDrinking = make(chan interface{})
)

var workHistory sync.Map

func main() {
	// register food and drink
	registerFood()
	registerDrinking()

	// Register employees
	registerEmployees()

	// Initialize chefs, bartenders, and waiters
	initializeEmployees()

	// Customers order food and drinks
	placeOrder()

	// Fan-in pattern: Collect results from readyFood and readyDrinking into announcement
	announcement := fanIn(readyFood, readyDrinking)

	var foodOrderList = make(chan order.Order, len(manager.FoodLists))
	var drinkOrderList = make(chan order.Order, len(manager.DrinkLists))

	// fan out pattern: Distribute food and drink orders to chefs and bartenders
	working(foodOrderList, drinkOrderList)

	for _, order := range manager.FoodLists {
		wgFood.Add(1)
		foodOrderList <- order
	}
	close(foodOrderList)
	for _, order := range manager.DrinkLists {
		wgDrinking.Add(1)
		drinkOrderList <- order
	}
	close(drinkOrderList)

	// Wait for chefs and bartenders to finish
	go func() {
		wgFood.Wait()
		close(readyFood)
	}()
	go func() {
		wgDrinking.Wait()
		close(readyDrinking)
	}()

	for x := range announcement {
		val := x.(order.Order)
		if foodItem, ok := val.Item.(food.IFood); ok {
			log.Printf("Customer: %s is served %s", val.NameCustomer, foodItem.GetFoodName())
			continue
		}
		if drinkItem, ok := val.Item.(drinking.IDrinking); ok {
			log.Printf("Customer: %s is served %s", val.NameCustomer, drinkItem.GetDrinkingName())
			continue
		}
	}

	fmt.Println("All waiters finished serving!!!")

	workHistoryTrace()
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
}

func working(foodOrderList, drinkOrderList chan order.Order) {
	for _, chef := range chefs {
		chef.Ready <- employee.READY
		go chef.Work(readyFood, &wgFood, foodOrderList, &workHistory)
	}

	for _, bartender := range bartenders {
		bartender.Ready <- employee.READY
		go bartender.Work(readyDrinking, &wgDrinking, drinkOrderList, &workHistory)
	}
}

func registerEmployees() {
	employee.RegisterEmployee(employee.CHEF, func(index int) employee.IEmployee {
		return employee.NewChef(index)
	})
	employee.RegisterEmployee(employee.BARTENDER, func(index int) employee.IEmployee {
		return employee.NewBartender(index)
	})
}

func registerFood() {
	food.RegisterFood(food.PIZZA, func() food.IFood {
		return food.NewPizza()
	})
	food.RegisterFood(food.BURGER, func() food.IFood {
		return food.NewBurger()
	})
	food.RegisterFood(food.PASTA, func() food.IFood {
		return food.NewPasta()
	})
}

func registerDrinking() {
	drinking.RegisterDrinking(drinking.JUICE, func() drinking.IDrinking {
		return drinking.NewJuice()
	})
	drinking.RegisterDrinking(drinking.TEA, func() drinking.IDrinking {
		return drinking.NewTea()
	})
	drinking.RegisterDrinking(drinking.COFFEE, func() drinking.IDrinking {
		return drinking.NewCoffee()
	})
}

func placeOrder() {
	Luffy.Order(food.PIZZA)
	Luffy.Order(drinking.JUICE)
	Luffy.Order(drinking.TEA)

	Zoro.Order(drinking.COFFEE)
	Zoro.Order(food.BURGER)
	Zoro.Order(food.BURGER)

	Chopper.Order(food.PIZZA)
	Chopper.Order(food.BURGER)
	Chopper.Order(drinking.JUICE)
}

func fanIn(readyFood, readyDrinking <-chan interface{}) <-chan interface{} {
	announcement := make(chan interface{})
	go func() {
		defer close(announcement)
		for val := range manager.Listen(readyFood, readyDrinking) {
			announcement <- val
		}
	}()
	return announcement
}

func workHistoryTrace() {
	fmt.Println("Working history:")
	workHistory.Range(func(key, count interface{}) bool {
		if employee, ok := key.(*employee.Chef); ok {
			fmt.Printf("Chef %d worked on %v orders\n", employee.Employee.ID, count)
		}
		if employee, ok := key.(*employee.Bartender); ok {
			fmt.Printf("Bartender %d worked on %v orders\n", employee.Employee.ID, count)
		}
		return true
	})
}
