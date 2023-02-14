package main

// concrete command

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}
