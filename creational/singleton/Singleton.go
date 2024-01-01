package singleton

import "fmt"

func TestSingleton() {
	for i := 0; i < 10; i++ {
		go GetInstance()
	}
	fmt.Scanln()
}
