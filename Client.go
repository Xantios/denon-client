package client

import (
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"net/http"
	"strings"
)

type Client struct {
	Host string
}

func New(host string) *Client {
	return &Client{Host: host}
}

func (c *Client) Request(name string) (*xmlquery.Node, error) {
	body := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<tx>\n  <cmd id=\"1\">%s</cmd>\n</tx>", name)
	reader := strings.NewReader(body)
	resp, err := http.Post(fmt.Sprintf("http://%s/goform/AppCommand.xml", c.Host), "application/xml", reader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid status code: " + resp.Status)
	}

	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
