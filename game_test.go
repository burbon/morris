package morris

import (
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

func TestPlayerManagerNext(t *testing.T) {
}

func TestGame(t *testing.T) {
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

	t.Run("next player", func(t *testing.T) {
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
	})

	t.Run("board empty on start", func(t *testing.T) {
	})

}
