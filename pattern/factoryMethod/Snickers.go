package main

type Snickers struct {
	CandyBar
}

func newSnickers() iCandyBar {
	return &Snickers{CandyBar{
		Name:     "Snickers",
		Price:    63.88,
		Quantity: 5,
		Calories: 243,
	}}
}
