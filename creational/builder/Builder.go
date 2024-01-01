package builder

import "fmt"

func TestBuilder() {
	d := NewDirector(NewNormalBuilder())
	fmt.Println("house", d.buildHouse())

	d.setBuilder(NewLuxeryBuilder())
	fmt.Println("house", d.buildHouse())
}
