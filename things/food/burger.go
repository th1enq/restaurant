package food

import "restaurant/helper"

type Burger struct {
	Food
}

func NewBurger() *Burger {
	return &Burger{Food{status: PENDING}}
}

func (b *Burger) GetName() string {
	return BURGER
}

func (b *Burger) GetStatus() string {
	return b.status
}

func (b *Burger) SetStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}
