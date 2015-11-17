package game

import "testing"

type entryCheck struct {
	X, Y          byte
	ExpectedPiece Entry
}

func TestNewStandardGameLayout(t *testing.T) {
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

func TestGetLocationOutsideBoardGivesError(t *testing.T) {
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

func TestSetLocationOutsideBoardGivesError(t *testing.T) {
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

func TestCompareTwoStandardGames(t *testing.T) {
	game1 := NewStandardGame()
	game2 := NewStandardGame()
	result, explanation := Compare(game1, game2)
	if result != true {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation != "" {
		t.Error("Expected an empty error, got ", explanation)
	}
}

func TestCompareTwoEmptyGames(t *testing.T) {
	game1, _ := SetupEmptyGame()
	game2, _ := SetupEmptyGame()
	result, explanation := Compare(game1, game2)
	if result != true {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation != "" {
		t.Error("Expected an empty error, got ", explanation)
	}
}

func TestCompareTwoCenteredGames(t *testing.T) {
	game1, _ := SetupCenteredGame()
	game2, _ := SetupCenteredGame()
	result, explanation := Compare(game1, game2)
	if result != true {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation != "" {
		t.Error("Expected an empty error, got ", explanation)
	}
}

func TestCompareTwoSquaredGames(t *testing.T) {
	game1, _ := SetupSquaredGame()
	game2, _ := SetupSquaredGame()
	result, explanation := Compare(game1, game2)
	if result != true {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation != "" {
		t.Error("Expected an empty error, got ", explanation)
	}
}

func TestCompareSameGame(t *testing.T) {
	game1 := NewStandardGame()
	result, explanation := Compare(game1, game1)
	if result != true {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation != "" {
		t.Error("Expected an empty error, got ", explanation)
	}
}

func TestCompareDifferentState(t *testing.T) {
	game1 := NewStandardGame()
	game2 := NewStandardGame()
	game1.State = Player2Win
	result, explanation := Compare(game1, game2)
	if result != false {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation == "" {
		t.Error("Expected an explanation, got nothing")
	}
}

func TestCompareDifferentInOnePiece(t *testing.T) {
	game1 := NewStandardGame()
	game2 := NewStandardGame()
	game1.SetLocation(1, 1, Player1)
	result, explanation := Compare(game1, game2)
	if result != false {
		t.Error("Expected the two games to be the same but got ", result)
	}
	if explanation == "" {
		t.Error("Expected an explanation, got nothing")
	}
}
