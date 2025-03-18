package food

import (
	"restaurant/helper"
)

const (
	PENDING = "pending"
	MAKING  = "making"
	SERVED  = "served"
)

const (
	PIZZA  = "pizza"
	PASTA  = "pasta"
	BURGER = "burger"
	NONE   = "none"
)

type Food struct {
	status string
	recipe []string
}

type IFood interface {
	GetFoodName() string
	GetFoodStatus() string
	SetFoodStatus(string) error
	GetFoodStep() []string
}

func (b *Food) GetFoodStatus() string {
	return b.status
}

func (b *Food) SetFoodStatus(status string) error {
	if status != PENDING && status != MAKING && status != SERVED {
		return helper.ErrInvalidStatus
	}
	b.status = status
	return nil
}

func (b *Food) GetFoodStep() []string {
	return b.recipe
}

func GetFood(name string) (IFood, error) {
	switch name {
	case PIZZA:
		return NewPizza(), nil
	case PASTA:
		return NewPasta(), nil
	case BURGER:
		return NewBurger(), nil
	default:
		return nil, helper.ErrInvalidFood
	}
}
