package adapter

type Client struct{}

func (c *Client) ChargeComputer(computer IAdapter) {
	computer.ConnectUSBPort()
}
