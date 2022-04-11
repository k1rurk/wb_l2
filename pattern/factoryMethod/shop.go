package main

func getCandyBar(candyType string) iCandyBar {
	switch candyType {
	case "twix":
		return newTwix()
	case "bounty":
		return newBounty()
	case "snickers":
		return newSnickers()
	default:
		return newTwix()
	}
}
