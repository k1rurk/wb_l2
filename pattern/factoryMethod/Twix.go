package main

type Twix struct {
	CandyBar
}

func newTwix() iCandyBar {
	return &Twix{CandyBar{
		Name:     "Twix",
		Price:    50.3,
		Quantity: 4,
		Calories: 250,
	}}
}
