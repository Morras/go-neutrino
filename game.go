package neutrino

import "fmt"
import "math"

type GameController struct{
	game *Game

	moveChannel chan Move
	stateChannel chan State
}

func (self *GameController) StartGame(g *Game) (<-chan Move, <-chan State) {
	self.game = g
	self.moveChannel = make(chan Move)
	self.stateChannel = make(chan State)
	return self.moveChannel, self.stateChannel
}

func (self *GameController) MakeMove(m Move) (State, error) {
	legalMove, errorMessage := self.isMoveLegal(m)
	if !legalMove {
		return self.game.State, fmt.Errorf(errorMessage)
	}

	self.move(m)

	winnerExists, winnerState := self.isThereAWinner()
	if winnerExists {
		self.stateChannel <- winnerState
		return winnerState, nil
	}

	self.game.State = self.getNextState()
	self.stateChannel <- self.game.State
	return self.game.State, nil
}

func (self *GameController) EndGame() {
	close(self.moveChannel)
	close(self.stateChannel)
}

func (self *GameController) isMoveLegal(move Move) (bool, string) {

	answer, errorMessage := self.isMoveValidForState(move)

	if ! answer {
		return false, errorMessage
	}

	deltaX := int8(move.ToX - move.FromX);
	deltaY := int8(move.ToY - move.FromY);

	if deltaX != 0 && deltaY != 0 && ( deltaX != deltaY && deltaX != -deltaY) {
		return false, "Piece must be move in a straight line"
	}

	direction := ""
	if deltaY < 0 {
		direction += "N"
	} else if deltaY > 0 {
		direction += "S"
	}
	if deltaX < 0 {
		direction += "W"
	} else if deltaX > 0 {
		direction += "E"
	}

	steps := math.Max(math.Abs(float64(deltaX)), math.Abs(float64(deltaY)))

	return self.isMoveByDirectionLegal(move.FromX, move.FromY, direction, byte(steps));
}

func (self *GameController) isMoveValidForState(move Move) (bool, string) {
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
		return false, "Move must start at board location with a player 1 piece"
	} else if entry == Player2 && state != Player2Move {
		return false, "Move must start at board location with a player 2 piece"
	} else if entry == Neutrino && state != Player1NeutrinoMove && state != Player2NeutrinoMove  {
		return false, "Move must start at board location with a player 1 piece"
	}
	return true, ""
}


func (self *GameController) isMoveByDirectionLegal(startX, startY byte, direction string, steps byte) (bool, string) {

	if steps == 0 || direction == "" {
		return false, fmt.Sprint("The suggested move does not actually move any piece.")
	}

	isNextSquareFree := true

	for i := byte(1); i <= steps; i++ {
		switch direction {
		case "N":
			free, errorString := self.checkIfSquareIsFree(startX - i, startY)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX - (i + 1), startY)
			if ! free {
				return free, errorString
			}
		case "NE":
			free, errorString := self.checkIfSquareIsFree(startX - i, startY + i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX - (i + 1), startY + (i + 1))
			if ! free {
				return free, errorString
			}
		case "E":
			free, errorString := self.checkIfSquareIsFree(startX, startY + i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX, startY + (i + 1))
			if ! free {
				return free, errorString
			}
		case "SE":
			free, errorString := self.checkIfSquareIsFree(startX + i, startY + i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX + (i + 1), startY + (i + 1))
			if ! free {
				return free, errorString
			}
		case "S":
			free, errorString := self.checkIfSquareIsFree(startX + i, startY)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX + (i + 1), startY)
			if ! free {
				return free, errorString
			}
		case "SW":
			free, errorString := self.checkIfSquareIsFree(startX + i, startY - i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX + (i + 1), startY - (i + 1))
			if ! free {
				return free, errorString
			}
		case "W":
			free, errorString := self.checkIfSquareIsFree(startX, startY - i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX, startY - (i + 1))
			if ! free {
				return free, errorString
			}
		case "NW":
			free, errorString := self.checkIfSquareIsFree(startX - i, startY - i)
			isNextSquareFree, _ = self.checkIfSquareIsFree(startX - (i + 1), startY - (i + 1))
			if ! free {
				return free, errorString
			}
		}
	}

	if isNextSquareFree {
		return false, "Move does not move untill an obstacle is hit"
	}

	return true, "no error"
}

func (self *GameController) checkIfSquareIsFree(x, y byte) (bool, string){
	entry, err := self.game.GetLocation(x, y)
	if err != nil {
		return false, err.Error()
	} else if entry != EmptySquare {
		return false, "Invalid move, cannot pass another piece"
	}
	return true, "No error"
}


func (self *GameController) move(move Move) {
	var newEntry Entry

	switch self.game.State {
	case Player1NeutrinoMove:
		newEntry = Neutrino
	case Player2NeutrinoMove:
		newEntry = Neutrino
	case Player1Move:
		newEntry = Player1
	case Player2Move:
		newEntry = Player2
	default:
		panic(fmt.Sprintf("Unable to perform move at the current state %d", self.game.State))
	}

	self.game.SetLocation(move.FromX, move.FromY, EmptySquare)
	self.game.SetLocation(move.ToX, move.ToY, newEntry)
	self.moveChannel <- move
}

func (self *GameController) isThereAWinner() (bool, State) {
	//TODO
	//See if neutrino is on back row
	//See if neutrino is on front row
	//See if neutrino is blocked
	//Calculate who is winner if it is blocked
	return false, self.game.State
}

func (self *GameController) getNextState() State {
	switch self.game.State {
	case Player1NeutrinoMove:
		return Player1Move
	case Player1Move:
		return Player2NeutrinoMove
	case Player2NeutrinoMove:
		return Player2Move
	case Player2Move:
		return Player1NeutrinoMove
	}
	//We should never get here
	panic(fmt.Sprintf("Game is in a state it cannot move on from %d", self.game.State))
}
