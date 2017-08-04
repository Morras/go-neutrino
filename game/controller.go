package game

import "fmt"
import "math"

type Controller struct {
	game *Game
}

func (self *Controller) PlayGame(g *Game) {
	self.game = g
}

func (self *Controller) MakeMove(m Move) (State, error) {

	if legalMove, errorMsg := self.isMoveLegal(m); !legalMove {
		return self.game.State, fmt.Errorf(errorMsg)
	}

	self.move(m)

	winnerExists, winnerState, err := self.isThereAWinner()
	if err != nil {
		//Reset the move so that both the board and the state
		//reflects that nothing has happened due to an error.
		//
		//The error is that no neutrino can be found and
		//the game cannot be played at this point.
		movedEntry, _ := self.game.GetLocation(m.ToX, m.ToY)
		self.game.SetLocation(m.FromX, m.FromY, movedEntry)
		self.game.SetLocation(m.ToX, m.ToY, EmptySquare)
		return self.game.State, err
	}

	if winnerExists {
		self.game.State = winnerState
		return winnerState, nil
	}

	self.game.State = self.getNextState()
	return self.game.State, nil
}

func (self *Controller) isMoveLegal(move Move) (answer bool, message string) {

	if answer, msg := self.isMoveValidForState(move); !answer {
		return false, msg
	}

	if (move.ToY == 0 && move.FromY != 0 && self.game.State == Player1Move && self.getOwnPiecesOnHomeRow(1) == 4) ||
		(move.ToY == 4 && move.FromY != 4 && self.game.State == Player2Move && self.getOwnPiecesOnHomeRow(2) == 4) {
		return false, "Cannot move all five pieces back on home row"
	}

	//Need to change from byte to int8 to prevent underflow
	deltaX := int8(move.ToX - move.FromX)
	deltaY := int8(move.ToY - move.FromY)

	if deltaX != 0 && deltaY != 0 && (deltaX != deltaY && deltaX != -deltaY) {
		return false, "Piece must be move in a straight line"
	}

	direction := Origo
	if deltaY < 0 {
		direction += NorthOffset
	} else if deltaY > 0 {
		direction += SouthOffset
	}
	if deltaX < 0 {
		direction += WestOffset
	} else if deltaX > 0 {
		direction += EastOffset
	}

	steps := math.Max(math.Abs(float64(deltaX)), math.Abs(float64(deltaY)))

	return self.isMoveByDirectionLegal(move.FromX, move.FromY, direction, byte(steps))
}

func (self *Controller) isMoveValidForState(move Move) (answer bool, message string) {

	state := self.game.State
	if state == Player1Win || state == Player2Win {
		return false, "Cannot move as the game has been won"
	}

	entry, err := self.game.GetLocation(move.FromX, move.FromY)
	if err != nil {
		return false, err.Error()
	}

	if entry == EmptySquare {
		return false, "Move must start at a non empty board location"
	} else if entry == Player1 && state != Player1Move {
		return false, "It must be player 1 turn to move player 1 piece"
	} else if entry == Player2 && state != Player2Move {
		return false, "It must be player 2 turn to move player 2 piece"
	} else if entry == Neutrino && state != Player1NeutrinoMove && state != Player2NeutrinoMove {
		return false, "It must be either player 1 or player 2 turn to move neutrino"
	}
	return true, ""
}

func (self *Controller) isMoveByDirectionLegal(startX, startY byte, direction Direction, steps byte) (answer bool, message string) {

	if steps == 0 || direction == Origo {
		return false, "The suggested move does not actually move any piece."
	}

	for i := byte(1); i <= steps; i++ {
		//free :=
		if ! self.checkIfNthNeighbourIsFree(startX, startY, i, direction) {
			return false, "Invalid move, cannot pass another piece"
		}
	}
	isNextSquareFree := self.checkIfNthNeighbourIsFree(startX, startY, steps+1, direction)
	if isNextSquareFree {
		return false, "Move does not move untill an obstacle is hit"
	}

	return true, ""
}

func (self *Controller) checkIfNthNeighbourIsFree(startX, startY, n byte, direction Direction) bool {
	switch direction {
	case N:
		return self.checkIfSquareIsFree(startX, startY-n)
	case NE:
		return self.checkIfSquareIsFree(startX+n, startY-n)
	case E:
		return self.checkIfSquareIsFree(startX+n, startY)
	case SE:
		return self.checkIfSquareIsFree(startX+n, startY+n)
	case S:
		return self.checkIfSquareIsFree(startX, startY+n)
	case SW:
		return self.checkIfSquareIsFree(startX-n, startY+n)
	case W:
		return self.checkIfSquareIsFree(startX-n, startY)
	case NW:
		return self.checkIfSquareIsFree(startX-n, startY-n)
	default:
		//Something has gone horrible wrong in the calculations
		panic(fmt.Sprintf("Game is trying to make a move in an invalid direction", direction))
	}
}

func (self *Controller) checkIfSquareIsFree(x, y byte) bool {
	entry, err := self.game.GetLocation(x, y)
	if err != nil {
		return false
	} else if entry != EmptySquare {
		return false
	}
	return true
}

func (self *Controller) move(move Move) {
	newEntry, _ := self.game.GetLocation(move.FromX, move.FromY)
	self.game.SetLocation(move.FromX, move.FromY, EmptySquare)
	self.game.SetLocation(move.ToX, move.ToY, newEntry)
}

func (self *Controller) isThereAWinner() (bool, State, error) {

	//neutrino location
	x, y := self.locateNeutrino()
	if x == 99 || y == 99 {
		return false, self.game.State, ErrNoNeutrinoInGame
	}

	//Note that if the neutrino has been move
	//it cannot be blocked so the only two valid
	//states where a blocked neutrino could happen
	//are Player1Move and Player2Move w
	isSquareBlocked := self.isSquareBlocked(x, y)
	if isSquareBlocked && self.game.State == Player1Move {
		return true, Player1Win, nil
	} else if isSquareBlocked && self.game.State == Player2Move {
		return true, Player2Win, nil
	}

	if y == 0 {
		return true, Player2Win, nil
	} else if y == 4 {
		return true, Player1Win, nil
	}
	return false, self.game.State, nil
}

func (self *Controller) locateNeutrino() (byte, byte) {
	for x := byte(0); x < 5; x++ {
		for y := byte(0); y < 5; y++ {
			piece, _ := self.game.GetLocation(x, y)
			if piece == Neutrino {
				return x, y
			}
		}
	}
	return 99, 99
}

func (self *Controller) isSquareBlocked(x, y byte) bool {
	isSquareBlocked := true
	iStart := -1
	if x == 0 {
		iStart = 0
	}
	jStart := -1
	if y == 0 {
		jStart = 0
	}
	for i := iStart; i < 2; i++ {
		for j := jStart; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			piece, err := self.game.GetLocation(byte(int(x)+i), byte(int(y)+j))
			isNeighbourBlocked := err != nil || piece != EmptySquare
			isSquareBlocked = isSquareBlocked && isNeighbourBlocked
		}
	}
	return isSquareBlocked
}

func (self *Controller) getNextState() State {
	switch self.game.State {
	case Player1NeutrinoMove:
		return Player1Move
	case Player1Move:
		return Player2NeutrinoMove
	case Player2NeutrinoMove:
		return Player2Move
	case Player2Move:
		return Player1NeutrinoMove
	default:
		//We should never get here
		panic(fmt.Sprintf("Game is in a state it cannot move on from %d", self.game.State))
	}
}

func (self *Controller) getOwnPiecesOnHomeRow(player byte) byte {
	count := byte(0)
	if player == 1 {
		for i := byte(0); i < 5; i++ {
			piece, _ := self.game.GetLocation(i, 0)
			if piece == Player1 {
				count++
			}
		}
	} else if player == 2 {
		for i := byte(0); i < 5; i++ {
			piece, _ := self.game.GetLocation(i, 4)
			if piece == Player2 {
				count++
			}
		}
	}
	return count
}
