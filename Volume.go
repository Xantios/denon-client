package client

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"net/http"
	"strings"
)

func MapDecibelToPercentage(decibel float64) float64 {
	minDb := -79.0
	maxDb := -6.0

	// Normalize the decibel value to be in the range [0, 1]
	normalized := (decibel - minDb) / (maxDb - minDb)

	// Map the normalized decibel value to the percentage range [0, 100]
	return normalized * 100
}

func MapPercentageToDecibel(percentageInt int) float64 {

	percentage := float64(percentageInt)

	minDecibel := -79.0
	maxDecibel := -6.0

	// Normalize the percentage value to be in the range [0, 1]
	normalized := percentage / 100.0

	// Map the normalized percentage value to the decibel range
	return minDecibel + normalized*(maxDecibel-minDecibel)
}

func (c *Client) GetMute() bool {
	resp, err := c.Request("GetMuteStatus")
	if err != nil {
		// return false
	}

	status := strings.ToLower(xmlquery.FindOne(resp, "//mute").InnerText())
	if status == "on" {
		return true
	}

	return false
}

func (c *Client) VolumeUp() {
	http.Get(fmt.Sprintf("http://%s/goform/formiPhoneAppDirect.xml?MVUP", c.Host))
}

func (c *Client) VolumeDown() {
	http.Get(fmt.Sprintf("http://%s/goform/formiPhoneAppDirect.xml?MVDOWN", c.Host))
}

func (c *Client) VolumeSet(db float64) {
	http.Get(fmt.Sprintf("http://%s/goform/formiPhoneAppVolume.xml?1+-%.2f", c.Host, db))
}

func (c *Client) VolumeSetPercentage(volume int) {
	value := MapPercentageToDecibel(volume)
	c.VolumeSet(value)
}
