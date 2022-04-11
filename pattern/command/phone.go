package main

import "fmt"

type Phone struct {
	isOn,
	isFlashOn,
	isBlueOn,
	isWifiOn bool
}

func NewPhone() *Phone {
	return &Phone{
		isOn:      false,
		isFlashOn: false,
		isBlueOn:  false,
		isWifiOn:  false,
	}
}

func (p *Phone) turnFlashlight() {
	if !p.isOn {
		fmt.Println("Телефон выключен, включите его, пожалуйста!")
		return
	}

	if p.isFlashOn {
		fmt.Println("Выключаем фонарик")
		p.isFlashOn = false
	} else {
		fmt.Println("Включаем фонарик")
		p.isFlashOn = true
	}
}

func (p *Phone) turnWifi() {
	if !p.isOn {
		fmt.Println("Телефон выключен, включите его, пожалуйста!")
		return
	}

	if p.isWifiOn {
		fmt.Println("Выключаем wi-fi")
		p.isWifiOn = false
	} else {
		fmt.Println("Включаем wi-fi")
		p.isWifiOn = true
	}
}

func (p *Phone) turnBluetooth() {
	if !p.isOn {
		fmt.Println("Телефон выключен, включите его, пожалуйста!")
		return
	}

	if p.isBlueOn {
		fmt.Println("Выключаем bluetooth")
		p.isBlueOn = false
	} else {
		fmt.Println("Включаем bluetooth")
		p.isBlueOn = true
	}
}

func (p *Phone) turnPower() {

	if p.isOn {
		fmt.Println("Выключаем телефон")
		p.isOn = false
	} else {
		fmt.Println("Включаем телефон")
		p.isOn = true
	}
}
