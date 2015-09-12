package neutrino

import "fmt"

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
	//TODO
	return true, "no error"
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
