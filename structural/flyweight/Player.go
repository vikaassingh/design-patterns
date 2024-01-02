package flyweight

type Player struct {
	Name      string
	Dress     IDress
	Latitude  float32
	Longitude float32
}

func NewPlayer(playerType Flyweight) *Player {
	dress, err := GetDressFactoryInstance().GetDressByPlayerType(playerType)
	if err != nil {
		return nil
	}

	return &Player{Dress: dress}
}

func (p *Player) NewLocation(lat, long float32) {
	p.Latitude = lat
	p.Longitude = long
}
