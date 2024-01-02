package flyweight

type Game struct {
	Players        []*Player
	CounterPlayers []*Player
}

func NewGame() *Game {
	return &Game{
		Players:        make([]*Player, 0),
		CounterPlayers: make([]*Player, 0),
	}
}

func (g *Game) AddPlayer() {
	player := NewPlayer(FlyweightPlayerDress)
	g.Players = append(g.Players, player)
}

func (g *Game) AddCounterPlayer() {
	player := NewPlayer(FlyweightCounterPlayerDress)
	g.CounterPlayers = append(g.CounterPlayers, player)
}
