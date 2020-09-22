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
	u1 := "User1"
	u2 := "User2"

	bp := NewPlayer(u1, PLAYER_COLOR_BLACK)
	wp := NewPlayer(u2, PLAYER_COLOR_WHITE)

	t.Run("initial state", func(t *testing.T) {
		pm := NewPlayerManager(bp, wp)

		got := pm.Next()
		if got != wp {
			t.Errorf("expected %v, got %v", wp.Name(), got.Name())
		}
	})

	t.Run("middle state", func(t *testing.T) {
		pm := NewPlayerManager(bp, wp)
		pm.state.last = wp

		got := pm.Next()
		if got != bp {
			t.Errorf("expected %v, got %v", bp.Name(), got.Name())
		}
	})
}

func TestGame(t *testing.T) {
	u1 := "User1"
	u2 := "User2"

	g := NewGame(u1, u2)

	if g.pm.black.User() != u1 {
		t.Errorf("expected %v, got %v", u1, g.pm.black.User())
	}
	if g.pm.white.User() != u2 {
		t.Errorf("expected %v, got %v", u1, g.pm.white.User())
	}
}
