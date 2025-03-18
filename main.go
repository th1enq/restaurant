package main

import (
	"fmt"
	"restaurant/employee"
	"sync"
)

const (
	NUMBEROFCHEFS      = 3
	NUMBEROFBARTENDERS = 3
	NUMBEROFWAITERS    = 3
)

var (
	manager    = employee.GetManager()
	chefs      = make([]interface{}, NUMBEROFCHEFS)
	bartenders = make([]interface{}, NUMBEROFBARTENDERS)
	waiters    = make([]*employee.Waiter, NUMBEROFWAITERS)
)

func main() {
	// var wg sync.WaitGroup
	var readyFood = make(chan interface{})
	var readyDrinking = make(chan interface{})

	for i := 0; i < NUMBEROFCHEFS; i++ {
		chefs[i], _ = employee.GetEmployee(i+1, employee.CHEF)
	}

	for i := 0; i < NUMBEROFWAITERS; i++ {
		waiters[i] = employee.NewWaiter(i + 1)
	}

	for i := 0; i < NUMBEROFBARTENDERS; i++ {
		bartenders[i], _ = employee.GetEmployee(i+1, employee.BARTENDER)
	}

	var wg sync.WaitGroup

	for _, chef := range chefs {
		wg.Add(1)
		go func(x interface{}) {
			x.(employee.IEmployee).Work(readyFood, &wg, "pizza")
		}(chef)
	}

	// for _, bartender := range bartenders {
	// 	wg.Add(1)
	// 	go func(x interface{}) {
	// 		x.(employee.IEmployee).Work(readyDrinking, &wg, "tea")
	// 	}(bartender)
	// }

	go func() {
		wg.Wait()
		close(readyFood)
		close(readyDrinking)
	}()

	announcement := make(chan interface{})

	go func() {
		for val := range manager.Listen(readyFood, readyDrinking) {
			announcement <- val
		}
		close(announcement)
	}()

	var wg2 sync.WaitGroup
	for _, waiter := range waiters {
		wg2.Add(1)
		go func(w *employee.Waiter) {
			w.Work(announcement, &wg2)
		}(waiter)
	}

	// Đợi tất cả waiters hoàn thành
	wg2.Wait()
	fmt.Println("All waiters finished serving!")
}
