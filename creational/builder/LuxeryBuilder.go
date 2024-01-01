package builder

type LuxeryBuilder struct {
	floor      int
	doorType   string
	windowType string
}

func NewLuxeryBuilder() *LuxeryBuilder {
	return &LuxeryBuilder{}
}

func (n *LuxeryBuilder) setFloorNumber() {
	n.floor = 3
}

func (n *LuxeryBuilder) setDoorType() {
	n.doorType = "snow door"
}

func (n *LuxeryBuilder) setWindowType() {
	n.windowType = "snow window"
}

func (n *LuxeryBuilder) getHouse() House {
	return House{
		floor:      n.floor,
		doorType:   n.doorType,
		windowType: n.windowType,
	}
}
