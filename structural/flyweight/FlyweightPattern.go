package flyweight

import "fmt"

type Flyweight int

const (
	FlyweightPlayerDress Flyweight = iota
	FlyweightCounterPlayerDress
)

func TestFlyweight() {
	var game Game
	game.AddPlayer()
	game.AddPlayer()
	game.AddPlayer()
	game.AddPlayer()

	game.AddCounterPlayer()
	game.AddCounterPlayer()
	game.AddCounterPlayer()
	game.AddCounterPlayer()

	dressFactoryInstance := GetDressFactoryInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("\nPlayerType: %d & DressColor: %s", dressType, dress.GetDressColor())
	}
}
