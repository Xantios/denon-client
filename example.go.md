package main

import "github.com/xantios/denon-client"

func main() {

	c := client.New("10.0.61")
	if !c.GetPowerStatus() {
		c.SetPowerStatus(true)
	}

	c.GetChannels()
	c.SetChannel("TV")

	c.VolumeSetPercentage(33)

	c.GetMute()
}
