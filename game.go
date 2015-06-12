package neutrino

var game *Game
var state State

var moveChannel chan Move
var stateChannel chan State

func StartGame(g *Game) (<-chan Move, <-chan State){
	game = g
	moveChannel = make(chan Move)
	stateChannel = make(chan State)
	return moveChannel, stateChannel
}

func MakeMove(move Move) error{
	moveChannel <- move
	state++
	stateChannel <- state
	return nil
}

func EndGame(){
	close(moveChannel)
	close(stateChannel)
}
