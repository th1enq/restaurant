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
	// channel for ready food and ready drinking
	var readyFood = make(chan interface{})
	var readyDrinking = make(chan interface{})

	// Initialize chefs, bartenders, and waiters
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

	// chef goroutine
	for _, chef := range chefs {
		wg.Add(1)
		c := chef.(employee.IEmployee)
		go c.Work(readyFood, &wg, "pasta")
	}

	// bartender goroutine
	for _, bartender := range bartenders {
		wg.Add(1)
		b := bartender.(employee.IEmployee)
		go b.Work(readyDrinking, &wg, "tea")
	}

	go func() {
		wg.Wait()
		close(readyFood)
		close(readyDrinking)
	}()

	announcement := make(chan interface{})

	// fan in pattern (readyFood and readyDrinking) -> announcement
	go func() {
		for val := range manager.Listen(readyFood, readyDrinking) {
			announcement <- val
		}
		close(announcement)
	}()

	// fan out pattern (announcement) -> waiters
	// waiter goroutine
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
