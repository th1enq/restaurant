package things

import (
	"restaurant/helper"
	"restaurant/things/drinking"
	"restaurant/things/food"
)

type Things interface {
	GetName() string
	GetStatus() string
	SetStatus(string) error
}

func GetThings(name string) (Things, error) {
	switch name {
	case drinking.COFFEE:
		return drinking.NewCoffee(), nil
	case drinking.JUICE:
		return drinking.NewJuice(), nil
	case drinking.TEA:
		return drinking.NewTea(), nil
	case food.PIZZA:
		return food.NewPizza(), nil
	case food.PASTA:
		return food.NewPasta(), nil
	case food.BURGER:
		return food.NewBurger(), nil
	default:
		return nil, helper.ErrInvalidFood
	}
}
