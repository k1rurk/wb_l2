package main

type Command interface {
	execute()
}

type FlashlightCommand struct {
	phone *Phone
}

func (c *FlashlightCommand) execute() {
	c.phone.turnFlashlight()
}

type BluetoothCommand struct {
	phone *Phone
}

func (c *BluetoothCommand) execute() {
	c.phone.turnBluetooth()
}

type WifiCommand struct {
	phone *Phone
}

func (c *WifiCommand) execute() {
	c.phone.turnWifi()
}

type PowerCommand struct {
	phone *Phone
}

func (c *PowerCommand) execute() {
	c.phone.turnPower()
}
