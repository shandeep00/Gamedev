package Entities

type client struct {
	clientId int
	name string
	status int
}

type Client struct {
	*client
}

func (c *client) updateStatus(status int) {
	c.status = status
}

func NewClient(clientId int, name string, status int) *Client{
	c := &client{
		clientId: clientId,
		name: name,
		status: status,
	}
	C := &Client{c}
	return C
}