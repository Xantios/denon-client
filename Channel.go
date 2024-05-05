package client

import (
	"github.com/antchfx/xmlquery"
	"net/http"
	"strings"
)

func (c *Client) GetCurrentChannel() string {
	resp, err := c.Request("GetSourceStatus")
	if err != nil {
		return "UNKNOWN"
	}

	channel := xmlquery.FindOne(resp, "//source").InnerText()
	return channel
}

func (c *Client) GetChannels() []string {
	resp, err := c.Request("GetRenameSource")
	if err != nil {
		return nil
	}

	channels := xmlquery.Find(resp, "//list")
	output := make([]string, len(channels))
	for _, channel := range channels {
		channelName := xmlquery.FindOne(channel, "//name").InnerText()
		output = append(output, channelName)
	}

	return output
}

func (c *Client) SetChannel(channel string) {
	// SI is SetInput ????
	patchedChannel := strings.ToUpper(channel)
	http.Get("http://" + c.Host + "/goform/formiPhoneAppDirect.xml?SI" + patchedChannel)
}
