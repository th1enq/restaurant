package food

type Pizza struct {
	Food
}

func NewPizza() *Pizza {
	return &Pizza{
		Food{
			Recipe: []string{
				"Prepare the dough",
				"Add the sauce",
				"Add the toppings",
			},
		},
	}
}

func (b *Pizza) GetFoodName() string {
	return PIZZA
}
