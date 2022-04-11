package main

func main() {
	car := NewCar("Old engine", "Old tyre", "Just doors")
	airplane := NewAirplane("Old engine", "Old tyre", "Just doors")
	tractor := NewTractor("Old engine", "Old tyre", "Just doors")

	changeEngine := &ChangeEngine{}

	car.change(changeEngine)
	car.info()
	airplane.change(changeEngine)
	airplane.info()
	tractor.change(changeEngine)
	tractor.info()

	changeTyre := &ChangeTyre{}

	car.change(changeTyre)
	car.info()
	airplane.change(changeTyre)
	airplane.info()
	tractor.change(changeTyre)
	tractor.info()

}
