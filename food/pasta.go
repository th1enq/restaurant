package food

type Pasta struct {
	Food
}

func NewPasta() *Pasta {
	return &Pasta{
		Food{
			status: PENDING,
			recipe: map[int]string{
				1: "Boil the pasta",
				2: "Add the sauce",
				3: "Serve the pasta",
			},
		},
	}
}

func (b *Pasta) GetFoodName() string {
	return PASTA
}
