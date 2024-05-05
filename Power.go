package client

import (
	"github.com/antchfx/xmlquery"
	"net/http"
)

func (c *Client) GetPowerStatus() bool {
	resp, err := c.Request("GetAllZonePowerStatus")
	if err != nil {
		return false
	}

	status := xmlquery.FindOne(resp, "//zone1").InnerText()
	if status == "OFF" {
		return false
	}

	return true
}

func (c *Client) SetPowerStatus(status bool) {
	var power string
	if status == false {
		power = "Off"
	} else {
		power = "On"
	}

	http.Get("http://" + c.Host + "/goform/formiPhoneAppPower.xml?1+Power" + power)

	// GET to http://10.0.0.61/goform/formiPhoneAppPower.xml?1+PowerOn
	// Returns: <?xml version="1.0" encoding="utf-8" ?>
	//  <item>
	//    <Power>
	//      <value>
	//        ON
	//      </value>
	//    </Power>
	//  </item>
}
