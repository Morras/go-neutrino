package neutrino

import "testing"

type entryCheck struct {
	X, Y byte
	ExpectedPiece Entry
}

func TestNewStandardGameLayout(t* testing.T){
	game := NewStandardGame()

	expectedEntries := []entryCheck{
		//Player1 home row
		{0, 0, Player1},
		{1, 0, Player1},
		{2, 0, Player1},
		{3, 0, Player1},
		{4, 0, Player1},
		//Row 2
		{0, 1, EmptySquare},
		{1, 1, EmptySquare},
		{2, 1, EmptySquare},
		{3, 1, EmptySquare},
		{4, 1, EmptySquare},
		//Middle row
		{0, 2, EmptySquare},
		{1, 2, EmptySquare},
		{2, 2, Neutrino},
		{3, 2, EmptySquare},
		{4, 2, EmptySquare},
		//Row 4
		{0, 3, EmptySquare},
		{1, 3, EmptySquare},
		{2, 3, EmptySquare},
		{3, 3, EmptySquare},
		{4, 3, EmptySquare},
		//Player1 home row
		{0, 4, Player2},
		{1, 4, Player2},
		{2, 4, Player2},
		{3, 4, Player2},
		{4, 4, Player2}}

	for _, e := range expectedEntries {
		actualEntry, _ := game.GetLocation(e.X, e.Y)
		if actualEntry != e.ExpectedPiece {
			t.Errorf("Expected %q at (%d, %d) got %q", e.ExpectedPiece, e.X, e.Y, actualEntry)
		}
	}
}

func TestGetLocationOutsideBoardGivesError(t* testing.T){
	game := NewStandardGame()
	for i := byte(0); i < 10; i++ {
		for j := byte(0); j < 10; j++ {
			_, err := game.GetLocation(i, j)
			if i > 4 || j > 4 {
				if err == nil {
					t.Errorf("Expected to get an error when asking for location at (%d,  %d) but got no error", i, j)
				}
			} else {
				if err != nil {
					t.Errorf("Expected not to get an error when asking for location at (%d,  %d) but got %q", i, j, err)
				}
			}
		}
	}
}

func TestSetLocationOutsideBoardGivesError(t* testing.T){
	game := NewStandardGame()
	for i := byte(0); i < 10; i++ {
		for j := byte(0); j < 10; j++ {
			err := game.SetLocation(i, j, Player1)
			if i > 4 || j > 4 {
				if err == nil {
					t.Errorf("Expected to get an error when asking for location at (%d,  %d) but got no error", i, j)
				}
			} else {
				if err != nil {
					t.Errorf("Expected not to get an error when asking for location at (%d,  %d) but got %q", i, j, err)
				}
			}
		}
	}
}
