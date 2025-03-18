package food

type Burger struct {
	Food
}

func NewBurger() *Burger {
	return &Burger{
		Food{
			status: PENDING,
			recipe: map[int]string{
				1: "Prepare the patty",
				2: "Add the toppings",
				3: "Assemble the burger",
			},
		},
	}
}

func (b *Burger) GetFoodName() string {
	return BURGER
}
