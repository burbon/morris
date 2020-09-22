package morris

type PlayerColor int

func (pc PlayerColor) String() string {
	if pc == 0 {
		return "Black"
	} else if pc == 1 {
		return "White"
	} else {
		panic("Unknown color")
	}
}

const (
	PLAYER_COLOR_BLACK PlayerColor = 0
	PLAYER_COLOR_WHITE             = 1
)

const (
	PIECES_NO = 3
)

type Player struct {
	color  PlayerColor
	user   string
	pieces int
}

func NewPlayer(name string, color PlayerColor) *Player {
	return &Player{
		color,
		name,
		PIECES_NO,
	}
}

func (p Player) Name() string {
	return p.color.String()
}

func (p Player) User() string {
	return p.user
}

func (p Player) Pieces() int {
	return p.pieces
}

type State struct {
	last  *Player
	board [][]*PlayerColor
}

type PlayerManager struct {
	state State
	black *Player
	white *Player
}

func NewPlayerManager(black *Player, white *Player) *PlayerManager {
	return &PlayerManager{
		black: black,
		white: white,
	}
}

func (pm PlayerManager) Next() *Player {
	if pm.state.last == nil {
		return pm.white
	} else {
		if pm.state.last == pm.white {
			return pm.black
		} else if pm.state.last == pm.black {
			return pm.white
		} else {
			panic("player manager error")
		}
	}
}

func (pm PlayerManager) Move(i int, j int, pc *PlayerColor) {
}
