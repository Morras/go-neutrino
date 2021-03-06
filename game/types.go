package game

import (
	"errors"
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
* To allow arithmetic, Origo will be 25 and then the following rules apply:
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
	Player1NeutrinoMove State = iota
	Player1Move
	Player2NeutrinoMove
	Player2Move
	Player1Win
	Player2Win
)

//// entry type ////

type Entry byte

const (
	EmptySquare Entry = iota
	Player1
	Player2
	Neutrino
)
