package food

type Pizza struct {
	Food
}

func NewPizza() *Pizza {
	return &Pizza{
		Food{
			status: PENDING,
			recipe: []string{
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
