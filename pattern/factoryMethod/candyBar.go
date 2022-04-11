package main

type CandyBar struct {
	Name  string
	Price float32
	Quantity,
	Calories int
}

func (b *CandyBar) getName() string {
	return b.Name
}

func (b *CandyBar) getPrice() float32 {
	return b.Price
}

func (b *CandyBar) getQuantity() int {
	return b.Quantity
}

func (b *CandyBar) getCalories() int {
	return b.Calories
}
