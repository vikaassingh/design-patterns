package flyweight

import "fmt"

type DressFactory struct {
	DressMap map[Flyweight]IDress
}

var (
	DressFactoryInstance = &DressFactory{DressMap: make(map[Flyweight]IDress, 0)}
)

func GetDressFactoryInstance() *DressFactory {
	return DressFactoryInstance
}

func (d *DressFactory) GetDressByPlayerType(playerType Flyweight) (IDress, error) {
	if d.DressMap[playerType] != nil {
		return d.DressMap[playerType], nil
	}

	if playerType == FlyweightPlayerDress {
		d.DressMap[playerType] = NewPlayerDress()
		return d.DressMap[playerType], nil
	}

	if playerType == FlyweightCounterPlayerDress {
		d.DressMap[playerType] = NewCounterPlayerDress()
		return d.DressMap[playerType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}
