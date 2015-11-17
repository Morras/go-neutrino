package game

import "fmt"

//// Game type ////

type Game struct {
	game  [25]Entry
	State State
}

func Compare(a, b *Game) (bool, string) {
	if &a == &b {
		return true, ""
	}
	for x := byte(1); x < 5; x++ {
		for y := byte(1); y < 5; y++ {
			entryA, _ := a.GetLocation(x, y)
			entryB, _ := b.GetLocation(x, y)
			if entryA != entryB {
				return false, fmt.Sprintf("Difference in (%d, %d): %d and %d", x, y, entryA, entryB)
			}
		}
	}
	if a.State != b.State {
		return false, fmt.Sprintf("Different states %d and %d", a.State, b.State)
	}
	return true, ""
}

const (
	widthOfBoard = 5
)

func NewEmptyGame() *Game {
	game := &Game{}
	for x := byte(0); x < 5; x++ {
		for y := byte(0); y < 5; y++ {
			game.SetLocation(x, y, EmptySquare)
		}
	}
	game.State = Player1NeutrinoMove
	return game
}

func NewStandardGame() *Game {
	game := NewEmptyGame()
	for i := byte(0); i < 5; i++ {
		game.SetLocation(i, 0, Player1)
		game.SetLocation(i, 4, Player2)
	}
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1NeutrinoMove
	return game
}

func (self *Game) SetLocation(x, y byte, entry Entry) error {
	if x > 4 || y > 4 {
		return fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	self.game[x+widthOfBoard*y] = entry
	return nil
}

func (self *Game) GetLocation(x, y byte) (Entry, error) {
	if x > 4 || y > 4 {
		return 9, fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	return self.game[x+widthOfBoard*y], nil
}
