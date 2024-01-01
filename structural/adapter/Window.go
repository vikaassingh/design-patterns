package adapter

import "fmt"

type Window struct{}

func (w *Window) ConnectUSBPort() {
	fmt.Println("Window Laptop is charging...")
}
