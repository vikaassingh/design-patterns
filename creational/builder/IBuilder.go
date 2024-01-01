package builder

type IBuilder interface {
	setDoorType()
	setWindowType()
	setFloorNumber()
	getHouse() House
}
