package main

type Bounty struct {
	CandyBar
}

func newBounty() iCandyBar {
	return &Bounty{CandyBar{
		Name:     "Bounty",
		Price:    44.90,
		Quantity: 9,
		Calories: 276,
	}}
}
