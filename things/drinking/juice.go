package drinking

import "restaurant/helper"

type Juice struct {
	Drinking
}

func NewJuice() *Juice {
	return &Juice{Drinking{status: PENDING}}
}

func (b *Juice) GetName() string {
	return JUICE
}

func (b *Juice) GetStatus() string {
	return b.status
}

func (b *Juice) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
