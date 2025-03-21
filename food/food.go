package food

import "restaurant/helper"

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

type FoodFactory func() IFood

var foodFactories = make(map[string]FoodFactory)

func RegisterFood(name string, factory FoodFactory) {
	foodFactories[name] = factory
}

func GetFood(name string) (IFood, error) {
	if factory, ok := foodFactories[name]; ok {
		return factory(), nil
	}
	return nil, helper.ErrInvalidFood
}
