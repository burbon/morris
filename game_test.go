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
					if g.board[x][y] != nil {
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
						if g.board[x][y] == nil || (*g.board[x][y] != currentPlayer.color) {
							t.Errorf("expeced piece from next player %v, got %v", currentPlayer.color, g.board[x][y])
						}
					} else {
						if g.board[x][y] != nil {
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
