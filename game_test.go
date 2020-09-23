package morris

import (
	"fmt"
	"testing"
)

// TestPlayer tests player
func TestPlayer(t *testing.T) {
	u := "User"
	p := NewPlayer(u, PLAYER_COLOR_BLACK)

	t.Run("color black", func(t *testing.T) {
		if p.color != PLAYER_COLOR_BLACK {
			t.Errorf("expected PLAYER_COLOR_BLACK, got %v", p.color)
		}
	})

	t.Run("name", func(t *testing.T) {
		if p.Name() != "Black" {
			t.Errorf("expected Black, got %s", p.Name())
		}
	})

	t.Run("pieces", func(t *testing.T) {
		if p.Pieces() != PIECES_NO {
			t.Errorf("expected pieces, got %v", p.Pieces())
		}
	})
}

// TestGameInit tests game initializtion
func TestGameInit(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	t.Run("assigning colors to users", func(t *testing.T) {
		if g.black.User() != u1 {
			t.Errorf("expected %v, got %v", u1, g.black.User())
		}
		if g.white.User() != u2 {
			t.Errorf("expected %v, got %v", u1, g.white.User())
		}
	})

	t.Run("board empty on start", func(t *testing.T) {
		for y := range [BOARD_Y]int{} {
			for x := range [BOARD_X]int{} {
				t.Run("middle state", func(t *testing.T) {
					if g.board[x][y] != PLAYER_COLOR_UNSET {
						t.Errorf("board should be nil everywhere, got %v", g.board[x][y])
					}
				})
			}
		}
	})
}

// TestGameNext tests returning next player
func TestGameNext(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	t.Run("initial state", func(t *testing.T) {
		got := g.Next()
		if got != g.white {
			t.Errorf("expected %v, got %v", g.white.Name(), got.Name())
		}
	})

	t.Run("middle state", func(t *testing.T) {
		g.last = g.white

		got := g.Next()
		if got != g.black {
			t.Errorf("expected %v, got %v", g.black.Name(), got.Name())
		}
	})
}

func TestGamePlay(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)
	var source *Coords
	destination := Coords{1, 1}
	currentPlayer := g.Next()
	g.Play(source, destination)

	t.Run("board with first move", func(t *testing.T) {
		for y := range [BOARD_Y]int{} {
			for x := range [BOARD_X]int{} {
				t.Run(fmt.Sprintf("%v %v", x, y), func(t *testing.T) {
					if x == destination.x && y == destination.y {
						if g.board[x][y] == PLAYER_COLOR_UNSET || (g.board[x][y] != currentPlayer.color) {
							t.Errorf("expeced piece from next player %v, got %v", currentPlayer.color, g.board[x][y])
						}
					} else {
						if g.board[x][y] != PLAYER_COLOR_UNSET {
							t.Errorf("board should be nil everywhere else, got %v at [%v, %v]", g.board[x][y], x, y)
						}
					}
				})
			}
		}
	})

	t.Run("next person switch", func(t *testing.T) {
		got := g.Next()
		if got != g.black {
			t.Errorf("expected %v, got %v", g.black.Name(), got.Name())
		}
	})
}

func TestGameIsFinished(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	t.Run("empty board", func(t *testing.T) {
		if g.IsFinished() != false {
			t.Errorf("expected game to continue")
		}
	})

	t.Run("horizontal win", func(t *testing.T) {
		g.board = [BOARD_X][BOARD_Y]PlayerColor{
			[BOARD_Y]PlayerColor{PLAYER_COLOR_BLACK, PLAYER_COLOR_UNSET, PLAYER_COLOR_UNSET},
			[BOARD_Y]PlayerColor{PLAYER_COLOR_BLACK, PLAYER_COLOR_UNSET, PLAYER_COLOR_UNSET},
			[BOARD_Y]PlayerColor{PLAYER_COLOR_BLACK, PLAYER_COLOR_UNSET, PLAYER_COLOR_UNSET},
		}

		if g.IsFinished() != true {
			t.Errorf("expected end of game")
		}
	})

	t.Run("vertical win", func(t *testing.T) {
		g.board = [BOARD_X][BOARD_Y]PlayerColor{
			[BOARD_Y]PlayerColor{PLAYER_COLOR_BLACK, PLAYER_COLOR_BLACK, PLAYER_COLOR_BLACK},
			[BOARD_Y]PlayerColor{},
			[BOARD_Y]PlayerColor{},
		}

		if g.IsFinished() != true {
			t.Errorf("expected end of game")
		}
	})
	t.Run("diagonal win", func(t *testing.T) {
		g.board = [BOARD_X][BOARD_Y]PlayerColor{
			[BOARD_Y]PlayerColor{PLAYER_COLOR_BLACK, PLAYER_COLOR_UNSET, PLAYER_COLOR_UNSET},
			[BOARD_Y]PlayerColor{PLAYER_COLOR_UNSET, PLAYER_COLOR_BLACK, PLAYER_COLOR_UNSET},
			[BOARD_Y]PlayerColor{PLAYER_COLOR_UNSET, PLAYER_COLOR_UNSET, PLAYER_COLOR_BLACK},
		}

		if g.IsFinished() != true {
			t.Errorf("expected end of game")
		}
	})
}

type Move struct {
	src *Coords
	dst Coords
}

func assertGameScript(t *testing.T, g *Game, gameScript []Move) {
	t.Helper()

	for i, s := range gameScript {
		g.Play(s.src, s.dst)
		if i != len(gameScript)-1 {
			if g.IsFinished() != false {
				t.Errorf("expected game to continue")
			}
		} else {
			if g.IsFinished() != true {
				t.Errorf("expected end of game")
			}
			if g.Last().color != PLAYER_COLOR_WHITE {
				t.Errorf("expected winner white, got %v", g.Last().color)
			}
		}
	}
}

func TestGameRun(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	var noSource *Coords
	gameScript := []Move{
		Move{noSource, Coords{0, 0}},
		Move{noSource, Coords{1, 0}},
		Move{noSource, Coords{0, 1}},
		Move{noSource, Coords{1, 1}},
		Move{noSource, Coords{0, 2}},
	}

	assertGameScript(t, g, gameScript)
}

func TestGameRunAdvanced(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	var noSource *Coords
	gameScript := []Move{
		Move{noSource, Coords{0, 0}},
		Move{noSource, Coords{1, 0}},
		Move{noSource, Coords{0, 1}},
		Move{noSource, Coords{0, 2}},
		Move{noSource, Coords{1, 1}},
		Move{noSource, Coords{2, 1}},
		Move{&Coords{0, 1}, Coords{2, 2}},
	}

	assertGameScript(t, g, gameScript)
}
