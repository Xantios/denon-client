package client

type Status struct {
	Power bool `json:"power"`
	Mute  bool `json:"mute"`
}

func (c *Client) GetStatus() Status {
	return Status{}
}
