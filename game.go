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
	PIECES_NO                      = 3
	BOARD_X                        = 3
	BOARD_Y                        = 3
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

type Coords struct {
	x int
	y int
}

// Game holds game mechanic
type Game struct {
	black *Player
	white *Player
	last  *Player
	board [BOARD_X][BOARD_Y]*PlayerColor
}

func NewGame(black string, white string) *Game {
	bp := NewPlayer(black, PLAYER_COLOR_BLACK)
	wp := NewPlayer(white, PLAYER_COLOR_WHITE)

	board := [BOARD_X][BOARD_Y]*PlayerColor{
		[BOARD_Y]*PlayerColor{},
		[BOARD_Y]*PlayerColor{},
		[BOARD_Y]*PlayerColor{},
	}
	return &Game{bp, wp, nil, board}
}

// Next returns player whos next turn
func (g *Game) Next() *Player {
	if g.last == nil {
		return g.white
	} else {
		if g.last == g.white {
			return g.black
		} else if g.last == g.black {
			return g.white
		} else {
			panic("game error")
		}
	}
}

func (g *Game) Play(source *Coords, destination Coords) {
	p := g.Next()
	if source == nil {
		if p.pieces == 0 {
			panic("wrong move")
		}
		p.pieces -= 1
	} else {
		if source.x == destination.x && source.y == destination.y {
			panic("wrong move")
		}
		g.board[source.x][source.y] = nil
	}
	g.board[destination.x][destination.y] = &p.color
	g.last = p
}
