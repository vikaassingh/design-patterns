package adapter

import "fmt"

type Macbook struct{}

func (w *Macbook) ConnectLightningPort() {
	fmt.Println("Macbook Laptop is charging...")
}
