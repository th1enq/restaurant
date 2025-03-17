package employee

import "sync"

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

func (m *Manager) Listen(readyFood chan interface{}, readyDrinking chan interface{}) {
	// chRes := make(chan interface{})
	// for {
	// 	select {
	// 	case <-readyFood:
	// 		chRes <-
	// 	case <-readyDrinking:
	// 		// do something
	// 	}
	// }

}
