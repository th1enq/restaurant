package food

type Pizza struct {
	Food
}

func NewPizza() *Pizza {
	return &Pizza{
		Food{
			status: PENDING,
			recipe: map[int]string{
				1: "Step 1",
				2: "Step 2",
				3: "Step 3",
			},
		},
	}
}

func (b *Pizza) GetFoodName() string {
	return PIZZA
}
