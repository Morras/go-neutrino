package neutrino

import "fmt"

//// Move type ////

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

//// State type ////

type State byte

const (
	Player1NeutrinoMove State = 0
	Player1Move State = 1
	Player2NeutrinoMove State = 2
	Player2Move State = 3
	Player1Win State = 4
	Player2Win State = 5
)

//// entry type ////

type Entry byte

const (
	EmptySquare Entry = '0'
	Player1 Entry = '1'
	Player2 Entry = '2'
	Neutrino Entry = 'N'
)

//// Game type ////

type Game struct {
	game [25]Entry
	State State
}

const (
	widthOfBoard = 5;
)

func NewEmptyGame() *Game{
	game := &Game{}
	for x := byte(0); x < 5; x++ {
		for y := byte(0); y < 5; y++ {
			game.SetLocation(x, y, EmptySquare)
		}
	}
	game.State = Player1NeutrinoMove
	return game;
}

func NewStandardGame() *Game{
	game := NewEmptyGame()
	for i := byte(0); i < 5; i++ {
		game.SetLocation(i, 0, Player1)
		game.SetLocation(i, 4, Player2)
	}
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1NeutrinoMove
	return game
}

func (self *Game) SetLocation(x, y byte, entry Entry) error{
	if x > 4 || y > 4{
		return fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	self.game[x+widthOfBoard*y] = entry
	return nil
}

func (self *Game) GetLocation(x, y byte) (Entry, error){
	if x > 4 || y > 4{
		return 9, fmt.Errorf("Coordinates must be between (0,0) and (4,4) both inclusive. Was (%d, %d)", x, y)
	}
	return self.game[x+widthOfBoard*y], nil
}

/**
 * This has proven usefull for debugging.
 * Note that the entries are the
 * ascii value of the rune, so 48 = '0'
 */
func (self *Game) PrintGame(){
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Printf("%f", self.game[x+widthOfBoard*y])
		}
		fmt.Println()
	}
}
