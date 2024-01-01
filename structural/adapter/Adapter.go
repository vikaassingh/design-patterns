package adapter

import "fmt"

type Adapter struct {
	Window *Macbook
}

func (a *Adapter) ConnectUSBPort() {
	fmt.Println("USB connecting to Lightning port")
	a.Window.ConnectLightningPort()
}

func TestAdapter() {
	var client Client
	window := &Window{}
	macbook := &Macbook{}
	adapter := &Adapter{
		Window: macbook,
	}

	client.ChargeComputer(window)
	client.ChargeComputer(adapter)
}
