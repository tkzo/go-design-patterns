package main

// concrete command

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
