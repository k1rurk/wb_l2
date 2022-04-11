package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type CPU struct {
}

func (c *CPU) freeze() {
	fmt.Println("freeze")
}

func (c *CPU) jump(position int64) {
	fmt.Println("jump position =", position)
}

func (c *CPU) execute() {
	fmt.Println("execute")
}

type Memory struct {
}

func (m *Memory) load(position int64, data []byte) {
	fmt.Println("load position =", position, ", data =", data)
}

type HardDrive struct {
}

func (d *HardDrive) read(lba int64, size int) []byte {
	fmt.Println("read lba =", lba, ", size =", size)
	return make([]byte, size)
}

type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	return &Computer{
		cpu:       new(CPU),
		memory:    new(Memory),
		hardDrive: new(HardDrive),
	}
}

func (c *Computer) startComputer() {
	BootAddress := 1
	BootSector := 2
	SectorSize := 3
	c.cpu.freeze()
	c.memory.load(int64(BootAddress), c.hardDrive.read(int64(BootSector), SectorSize))
	c.cpu.execute()
}

func main() {
	comp := NewComputer()
	comp.startComputer()
}
