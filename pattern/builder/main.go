package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type IBurger interface {
	prepareBuns()
	cookMeat()
	addCheese()
	addIngredients()
	getBurger() Burger
}

type Burger struct {
	buns,
	meat,
	cheese string
	ingredients []string
}

func (b *Burger) info() {
	fmt.Println("Булочка:", b.buns)
	fmt.Println("Сыр:", b.cheese)
	fmt.Println("Котлета:", b.meat)
	fmt.Println("Ингрединеты: ")
	for _, v := range b.ingredients {
		fmt.Println(v)
	}
}

type FranceBurger struct {
	buns,
	meat,
	cheese string
	ingredients []string
}

func (b *FranceBurger) prepareBuns() {
	b.buns = "Булочка-бриошь"
}

func (b *FranceBurger) cookMeat() {
	b.meat = "Говяжий фарш 5-10% жирности"
}

func (b *FranceBurger) addCheese() {
	b.cheese = "Голубой сыр"
}

func (b *FranceBurger) addIngredients() {
	b.ingredients = append(b.ingredients, "Маслины", "Тархун")
}

func (b *FranceBurger) getBurger() Burger {
	return Burger{
		buns:        b.buns,
		meat:        b.meat,
		cheese:      b.cheese,
		ingredients: b.ingredients,
	}
}

type ItalianBurger struct {
	buns,
	meat,
	cheese string
	ingredients []string
}

func (b *ItalianBurger) prepareBuns() {
	b.buns = "Чиабатта"
}

func (b *ItalianBurger) cookMeat() {
	b.meat = "Ветчина"
}

func (b *ItalianBurger) addCheese() {
	b.cheese = "Моцарелла"
}

func (b *ItalianBurger) addIngredients() {
	b.ingredients = append(b.ingredients, "Песто", "Оливковое масло")
}

func (b *ItalianBurger) getBurger() Burger {
	return Burger{
		buns:        b.buns,
		meat:        b.meat,
		cheese:      b.cheese,
		ingredients: b.ingredients,
	}
}

type BurgerDirector struct {
	burger IBurger
}

func NewBurger(burger IBurger) *BurgerDirector {
	return &BurgerDirector{
		burger: burger,
	}
}

func (d *BurgerDirector) setBurger(burger IBurger) {
	d.burger = burger
}

func (d *BurgerDirector) buildBurger() Burger {
	d.burger.prepareBuns()
	d.burger.cookMeat()
	d.burger.addCheese()
	d.burger.addIngredients()
	return d.burger.getBurger()
}

func main() {
	italBurg := new(ItalianBurger)
	burgerDir := NewBurger(italBurg)
	burger := burgerDir.buildBurger()
	fmt.Println("Итальянский бургер")
	burger.info()
	fmt.Println()
	fmt.Println("Французский бургер")
	francBurger := new(FranceBurger)
	burgerDir.setBurger(francBurger)
	burger = burgerDir.buildBurger()
	burger.info()
}
