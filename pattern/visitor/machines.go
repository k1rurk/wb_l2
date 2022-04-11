package main

import "fmt"

type ChangeElement interface {
	change(ChangeVisitor)
}

type Car struct {
	Engine string
	Tyre   string
	Door   string
}

func NewCar(engine, tyre, door string) *Car {
	return &Car{
		Engine: engine,
		Tyre:   tyre,
		Door:   door,
	}
}

func (c *Car) change(visitor ChangeVisitor) {
	visitor.changeInCar(c)
}

func (c *Car) info() {
	fmt.Println("Car-------------")
	fmt.Println("Engine:", c.Engine)
	fmt.Println("Tyre:", c.Tyre)
	fmt.Println("Door:", c.Door)
	fmt.Println("----------------")
}

type Airplane struct {
	Engine string
	Tyre   string
	Door   string
}

func NewAirplane(engine, tyre, door string) *Airplane {
	return &Airplane{
		Engine: engine,
		Tyre:   tyre,
		Door:   door,
	}
}

func (a *Airplane) change(visitor ChangeVisitor) {
	visitor.changeInAirplane(a)
}

func (a *Airplane) info() {
	fmt.Println("Airplane--------")
	fmt.Println("Engine:", a.Engine)
	fmt.Println("Tyre:", a.Tyre)
	fmt.Println("Door:", a.Door)
	fmt.Println("----------------")
}

type Tractor struct {
	Engine string
	Tyre   string
	Door   string
}

func NewTractor(engine, tyre, door string) *Tractor {
	return &Tractor{
		Engine: engine,
		Tyre:   tyre,
		Door:   door,
	}
}

func (t *Tractor) change(visitor ChangeVisitor) {
	visitor.changeInTractor(t)
}

func (t *Tractor) info() {
	fmt.Println("Tractor---------")
	fmt.Println("Engine:", t.Engine)
	fmt.Println("Tyre:", t.Tyre)
	fmt.Println("Door:", t.Door)
	fmt.Println("----------------")
}
