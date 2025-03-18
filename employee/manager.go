package employee

import (
	"sync"
)

const NUMBEROFTHINGS = 2

type Manager struct {
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
		for {
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

			if readyFood == nil && readyDrinking == nil {
				return
			}
		}
	}()
	return announcement
}
