package main

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

func main() {
	phone := NewPhone()

	cWifi := &WifiCommand{phone: phone}
	cOn := &PowerCommand{phone: phone}
	cBlue := &BluetoothCommand{phone: phone}
	cFlash := &FlashlightCommand{phone: phone}

	flashButton := &Button{
		command: cFlash,
	}
	flashButton.press()

	powerButton := &Button{
		command: cOn,
	}

	powerButton.press()
	flashButton.press()
	flashButton.press()

	wifiButton := &Button{
		command: cWifi,
	}

	wifiButton.press()

	bluetoothButton := &Button{
		command: cBlue,
	}

	bluetoothButton.press()

}
