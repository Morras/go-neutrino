package neutrino

import "fmt"

type Move struct {
	FromX, FromY, ToX, ToY byte
}

func NewMove(fromX, fromY, toX, toY byte) Move {
	return Move{
		FromX: fromX,
		FromY: fromY,
		ToX: toX,
		ToY: toY,
	}
}

type State byte

const (
	Player1NeutrinoMove State = 0
	Player1Move State = 1
	Player2NeutrinoMove State = 2
	Player2Move State = 3
	Player1Win State = 4
	Player2Win State = 5
)

type entry byte

const (
	EmptySquare = '0'
	Player1 entry = '1'
	Player2 = '2'
	Neutrino = 'N'
)

type Game struct {
	game [25]entry
}

const (
	widthOfBoard = 5;
)

func NewGame() *Game{
	game := &Game{}
	for i := byte(0); i < 5; i++ {
		game.SetLocation(i, 0, Player1)
		game.SetLocation(i, 4, Player2)
	}
	game.SetLocation(2, 2, Neutrino)
	return game
}

func (self *Game) SetLocation(x, y byte, entry entry) error{
	if x > 4 || y > 4{
		return fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	self.game[x+widthOfBoard*y] = entry
	return nil
}

func (self *Game) GetLocation(x, y byte) (entry, error){
	if x > 4 || y > 4{
		return 9, fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	return self.game[x+widthOfBoard*y], nil
}
