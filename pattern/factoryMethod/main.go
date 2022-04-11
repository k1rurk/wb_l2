package main

import "fmt"

func main() {
	bounty := getCandyBar("bounty")
	snickers := getCandyBar("snickers")
	defaultCandy := getCandyBar("")

	candyInfo(bounty)
	candyInfo(snickers)
	candyInfo(defaultCandy)
}

func candyInfo(candy iCandyBar) {
	fmt.Println(candy.getName())
	fmt.Println(candy.getPrice())
	fmt.Println(candy.getQuantity())
	fmt.Println(candy.getCalories())
	fmt.Println()
}
