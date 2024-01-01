package builder

type NormalBuilder struct {
	floor      int
	doorType   string
	windowType string
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setFloorNumber() {
	n.floor = 1
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "wooden door"
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "wooden window"
}

func (n *NormalBuilder) getHouse() House {
	return House{
		floor:      n.floor,
		doorType:   n.doorType,
		windowType: n.windowType,
	}
}
