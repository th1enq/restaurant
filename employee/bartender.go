package employee

import (
	"fmt"
	"log"
	"math/rand"
	"restaurant/drinking"
	"restaurant/order"
	"sync"
	"time"
)

type Bartender struct {
	Employee
}

func (b *Bartender) Work(readyDrinking chan<- interface{}, wg *sync.WaitGroup, orderLists chan order.Order) {
	defer wg.Done()

	for orderItem := range orderLists {
		if b.Status != RELAX {
			continue
		}
		fmt.Println(b, orderItem)
		drinkItem, err := drinking.GetDrinking(orderItem.Item.(drinking.IDrinking).GetDrinkingName())
		if err != nil {
			panic(err)
		}
		b.SetStatus(WORKING)
		for _, step := range drinkItem.GetRecipe() {
			b.SetStatus(fmt.Sprintf("Bartender %d %s", b.ID, step))
			log.Println(b.Status)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		readyDrinking <- orderItem
		b.SetStatus(RELAX)
	}
}

func newBartender(id int) *Bartender {
	return &Bartender{
		Employee: Employee{
			Status: RELAX,
			ID:     id,
		},
	}
}
