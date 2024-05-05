# Denon API Client

## What ?
As Denon doesnt really give support on this AVR anymore 
i'd say it would be fair game to reverse engineer their mobile app

The mobile app is REALLY old and its a miracle it even runs on my modern iPhone

this proxy **DOES NOT** utilize the Telnet protocol !!

## How to figure this out?

#### Setup a proxy
Its just plain HTTP, just a simple proxy will work

#### Android stuff
Setup a android emulator: https://johnborg.es/2019/04/android-setup-macos.html
Use HTTPToolkit to pull requests out of the device

## Example 

```golang
package main

import "xantios.nl/denon-proxy/client"

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
```