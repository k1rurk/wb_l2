package main

type ChangeVisitor interface {
	changeInCar(*Car)
	changeInAirplane(*Airplane)
	changeInTractor(*Tractor)
}

type ChangeEngine struct {
}

func (e *ChangeEngine) changeInCar(car *Car) {
	car.Engine = "New engine"
}

func (e *ChangeEngine) changeInAirplane(airplane *Airplane) {
	airplane.Engine = "New engine"
}

func (e *ChangeEngine) changeInTractor(tractor *Tractor) {
	tractor.Engine = "New engine"
}

type ChangeTyre struct {
}

func (t *ChangeTyre) changeInCar(car *Car) {
	car.Tyre = "New tyre"
}

func (t *ChangeTyre) changeInAirplane(airplane *Airplane) {
	airplane.Tyre = "New tyre"
}

func (t *ChangeTyre) changeInTractor(tractor *Tractor) {
	tractor.Tyre = "New tyre"
}
