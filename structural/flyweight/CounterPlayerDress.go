package flyweight

type CounterPlayerDress struct {
	Color string
}

func NewCounterPlayerDress() IDress {
	return &PlayerDress{Color: "Black"}
}

func (p *CounterPlayerDress) GetDressColor() string {
	return p.Color
}
