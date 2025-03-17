package drinking

import "restaurant/helper"

type Coffee struct {
	Drinking
}

func NewCoffee() *Coffee {
	return &Coffee{Drinking{status: PENDING}}
}

func (b *Coffee) GetName() string {
	return COFFEE
}

func (b *Coffee) GetStatus() string {
	return b.status
}

func (b *Coffee) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
