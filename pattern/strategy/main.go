package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Compression interface {
	compress(string)
}

type Zip struct {
}

func (z *Zip) compress(filename string) {
	fmt.Println("Zip compression", filename)
}

type Rar struct {
}

func (r *Rar) compress(filename string) {
	fmt.Println("Rar compression", filename)
}

type ARJ struct {
}

func (a *ARJ) compress(filename string) {
	fmt.Println("ARJ compression", filename)
}

type Compressor struct {
	p Compression
}

func (c *Compressor) setStrategy(compression Compression) {
	c.p = compression
}

func (c *Compressor) compress(filename string) {
	c.p.compress(filename)
}

func main() {
	rar := new(Rar)
	comp := Compressor{p: rar}
	comp.compress("file.txt")
	zip := new(Zip)
	comp.setStrategy(zip)
	comp.compress("file2.txt")
}
