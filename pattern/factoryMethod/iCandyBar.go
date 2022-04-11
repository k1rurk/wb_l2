package main

type iCandyBar interface {
	getPrice() float32
	getQuantity() int
	getCalories() int
	getName() string
}
