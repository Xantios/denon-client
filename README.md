# Denon API Client

This should be compatible with the following devices:

    Denon AVR-X1000

    Denon AVR-X1100W

    Denon AVR-X1200W

    Denon AVR-X1300W

    Denon AVR-X1400H

    Denon AVR-X1500H

    Denon AVR-X1600H

    Denon AVR-X1700H

    Denon AVR-X1800H

    Denon AVR-X2000

    Denon AVR-X2100W

    Denon AVR-X2200W

    Denon AVR-X2300W

    Denon AVR-X2400H

    Denon AVR-X2500H

    Denon AVR-X2600H

    Denon AVR-X2700H

    Denon AVR-X2800H

    Denon AVR-X3000

    Denon AVR-X3200W

    Denon AVR-X3300W

    Denon AVR-X3400H

    Denon AVR-X3500H

    Denon AVR-X3600H

    Denon AVR-X3700H

    Denon AVR-X4100W

    Denon AVR-X4300H

    Denon AVR-X4400H

    Denon AVR-X4500H

    Denon AVR-X4700H

    Denon AVC-X4800H

    Denon AVR-X6500H

    Denon AVR-X6700H

    Denon AVR-X7200W

    Denon AVR-X8500H

    Denon AVR-1713

    Denon AVR-1912

    Denon AVR-2112CI

    Denon AVR-2312CI

    Denon AVR-3311CI

    Denon AVR-3312

    Denon AVR-3313CI

    Denon AVR-4810

    Denon AVR-S650H

    Denon AVR-S710W

    Denon AVR-S720W

    Denon AVR-S740H

    Denon AVR-S750H

    Denon AVR-S760H

    Denon AVR-S770H

    Denon AVR-S940H

    Denon AVR-S950H

    Denon AVR-S960H

    Denon DN-500AV

    Denon DRA-800H

    Marantz AV7702

    Marantz AV7703

    Marantz AV7704

    Marantz AV8802A

    Marantz CINEMA 50

    Marantz CINEMA 70s

    Marantz M-CR510

    Marantz M-CR511

    Marantz M-CR603

    Marantz M-CR610

    Marantz M-CR611

    Marantz SR5006

    Marantz SR5008

    Marantz SR5010

    Marantz SR5011

    Marantz SR5015

    Marantz SR6007 - SR6012

    Marantz SR7007

    Marantz SR7010

    Marantz SR7012

    Marantz SR8015

    Marantz NR1504

    Marantz NR1506

    Marantz NR1509

    Marantz NR1510

    Marantz NR1602

    Marantz NR1603

    Marantz NR1604

    Marantz NR1606

    Marantz NR1607

    Marantz NR1710

    Marantz NR1711

## What ?
As Denon doesnt really give support on these AVRs anymore 
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