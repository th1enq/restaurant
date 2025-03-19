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
	Recipe []string
}

type IFood interface {
	GetFoodName() string
	GetRecipe() []string
}

func (b *Food) GetRecipe() []string {
	return b.Recipe
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
