package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	On()
	Off()
	Print()
	AddPaper(count int)
}

type Printer struct {
	PaperOffState State
	PowerOffState State
	PrintState    State
	WaitingState  State

	CurrentState State

	countPaper int
}

func (p *Printer) SetState(state State) {
	p.CurrentState = state
}

func NewPrinter() *Printer {
	p := &Printer{
		countPaper: 0,
	}

	paperOffState := &PaperOffState{
		printer: p,
	}

	powerOffState := &PowerOffState{
		printer: p,
	}

	printState := &PrintState{
		printer: p,
	}

	waitingState := &WaitingState{
		printer: p,
	}

	p.SetState(waitingState)

	p.PaperOffState = paperOffState
	p.PrintState = printState
	p.WaitingState = waitingState
	p.PowerOffState = powerOffState

	return p
}

func (p *Printer) PrintDocument() {
	p.CurrentState.Print()
}

func (p *Printer) PowerOff() {
	p.CurrentState.Off()
}

func (p *Printer) PowerOn() {
	p.CurrentState.On()
}

func (p *Printer) AddPaper(count int) {
	p.countPaper += count
}

// PaperOffState Состояние отсутствия бумаги:
type PaperOffState struct {
	printer *Printer
}

func (s *PaperOffState) On() {
	fmt.Println("Принтер и так уже включен")
}

func (s *PaperOffState) Off() {
	fmt.Println("Принтер выключен")
	s.printer.SetState(s.printer.PowerOffState)
}

func (s *PaperOffState) Print() {
	if s.printer.countPaper > 0 {
		s.printer.SetState(s.printer.PrintState)
		s.printer.PrintDocument()
	} else {
		fmt.Println("Бумаги нет, печатать не буду")
	}
}

func (s *PaperOffState) AddPaper(count int) {
	if count > 0 {
		fmt.Println("Добавляем бумагу")
	}
	s.printer.AddPaper(count)
	if s.printer.countPaper > 0 {
		s.printer.SetState(s.printer.WaitingState)
	}
}

// PowerOffState Состояние отключённого питания:
type PowerOffState struct {
	printer *Printer
}

func (s *PowerOffState) On() {
	fmt.Println("Принтер включен")
	s.printer.SetState(s.printer.WaitingState)
}

func (s *PowerOffState) Off() {
	fmt.Println("Принтер и так выключен")
}

func (s *PowerOffState) Print() {
	fmt.Println("Принтер выключен, печать невозможна")
}

func (s *PowerOffState) AddPaper(count int) {
	s.printer.AddPaper(count)
	if count > 0 {
		fmt.Println("Добавляем бумагу")
	}
}

// PrintState Состояние печати:
type PrintState struct {
	printer *Printer
}

func (s *PrintState) On() {
	fmt.Println("Принтер и так включен")
}

func (s *PrintState) Off() {
	fmt.Println("Принтер отключен")
	s.printer.SetState(s.printer.PowerOffState)
}

func (s *PrintState) Print() {
	if s.printer.countPaper > 0 {
		fmt.Println("Идет печать...")
		s.AddPaper(-1)
		s.printer.SetState(s.printer.WaitingState)
	} else {
		s.printer.SetState(s.printer.PaperOffState)
		s.printer.PrintDocument()
	}
}

func (s *PrintState) AddPaper(count int) {
	s.printer.AddPaper(count)
	if count > 0 {
		fmt.Println("Добавляем бумагу")
	}
}

// WaitingState Состояние ожидания печати:
type WaitingState struct {
	printer *Printer
}

func (s *WaitingState) On() {
	fmt.Println("Принтер и так включен")
}

func (s *WaitingState) Off() {
	fmt.Println("Принтер выключен")
	s.printer.SetState(s.printer.PowerOffState)
}

func (s *WaitingState) Print() {
	if s.printer.countPaper > 0 {
		fmt.Println("Cейчас все распечатаем...")
		s.printer.AddPaper(-1)
	} else {
		s.printer.SetState(s.printer.PaperOffState)
		s.printer.PrintDocument()
	}
}

func (s *WaitingState) AddPaper(count int) {
	s.printer.AddPaper(count)
	if count > 0 {
		fmt.Println("Добавляем бумагу")
	}
}

func main() {
	p := NewPrinter()
	p.PowerOn()
	p.PrintDocument()
	p.AddPaper(3)
	p.PrintDocument()
	p.PrintDocument()
	p.PrintDocument()
	p.PrintDocument()
	p.PowerOff()
}
