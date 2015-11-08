package neutrino

import (
	"errors"
	"fmt"
)

//// Error type ////
var (
	ErrNoNeutrinoInGame = errors.New("Unable to locate neutrino")
)

//// Move type ////

type Move struct {
	FromX, FromY, ToX, ToY byte
}

func NewMove(fromX, fromY, toX, toY byte) Move {
	return Move{
		FromX: fromX,
		FromY: fromY,
		ToX:   toX,
		ToY:   toY,
	}
}

//// Direction type ////

type Direction int8

/**
* To allow arithmetic Origo weill be 25 and then the following rules apply:
, Moving left subtracts 1
* Moving right adds 1
* Moving up subtracts 10
* Moving down adds 10
*/
const (
	Origo Direction = 25
	NW    Direction = 14
	N     Direction = 15
	NE    Direction = 16
	E     Direction = 26
	SE    Direction = 36
	S     Direction = 35
	SW    Direction = 34
	W     Direction = 24
)

const (
	NorthOffset Direction = -10
	SouthOffset Direction = 10
	WestOffset  Direction = -1
	EastOffset  Direction = 1
)

//// State type ////

type State byte

const (
	Player1NeutrinoMove State = 0
	Player1Move         State = 1
	Player2NeutrinoMove State = 2
	Player2Move         State = 3
	Player1Win          State = 4
	Player2Win          State = 5
)

//// entry type ////

type Entry byte

const (
	EmptySquare Entry = 0
	Player1     Entry = 1
	Player2     Entry = 2
	Neutrino    Entry = 3
)

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
