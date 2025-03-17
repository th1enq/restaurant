package main

import (
	"fmt"
	"restaurant/employee"
	"restaurant/things"
	"restaurant/things/food"
)

const (
	numbersOfEmployee = 2
)

var (
	manager, _   = employee.GetEmployee(employee.MANAGER)
	chef, _      = employee.GetEmployee(employee.CHEF)
	bartender, _ = employee.GetEmployee(employee.BARTENDER)
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(numbersOfEmployee)

	// go chef.Work()
	// go bartender.Work()

	// manager.Listen(chef.GetReadyThing(), bartender.GetReadyThing())

	// wg.Wait()
	pizza, _ := things.GetThings(food.PIZZA)
	// if err != nil {
	// 	panic(err)
	// }

	t, errr := pizza.(*food.Pizza)

	if !errr {
		return
	}

	fmt.Println(t.GetName())
}
