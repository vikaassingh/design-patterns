package flyweight

type PlayerDress struct {
	Color string
}

func NewPlayerDress() IDress {
	return &PlayerDress{Color: "Red"}
}

func (p *PlayerDress) GetDressColor() string {
	return p.Color
}
