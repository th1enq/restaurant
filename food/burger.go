package food

type Burger struct {
	Food
}

func NewBurger() *Burger {
	return &Burger{
		Food{
			status: PENDING,
			recipe: []string{
				"Prepare the bun",
				"Grill the patty",
				"Serve the burger",
			},
		},
	}
}

func (b *Burger) GetFoodName() string {
	return BURGER
}
