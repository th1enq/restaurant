package food

type Burger struct {
	Food
}

func NewBurger() *Burger {
	return &Burger{
		Food{
			Recipe: []string{
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
