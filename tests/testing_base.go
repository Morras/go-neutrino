package neutrino

import . "github.com/morras/go-neutrino"

/** 
 * There is a problem that GameController sends messages
 * to a state and a move channel when a move is being 
 * made, and if these channels are not cleared the 
 * program deadlocks. 
 * This method makes sure to clear the channel and then
 * does nothing
 */
func pollChannels(moveChan <-chan Move, stateChan <-chan State) {
	go pollMoves(moveChan)
	go pollStates(stateChan)
}
func pollMoves(moveChan <-chan Move){
	for move := range moveChan{ move = move}
}

func pollStates(stateChan <-chan State){
	for state := range stateChan{ state = state}
}

/**
	* Basic setup of an empty game
	* Remember to add a neutrino to your game
	* otherwise it is invalid and the game controller
	* will panic
	*/
func setupEmptyGame() (*Game, *GameController) {
	game := NewEmptyGame()
	controller := &GameController{}
	mChan, sChan := controller.StartGame(game)
	go pollChannels(mChan, sChan)
	return game, controller
}

/**
 * Setup with a player 1 piece on
 * (1,1), (3,1), (3,1) and (3,3)
 * and it is player ones turn to move
 * the neutrino.
 */
func setupSquaredGame() (*Game, *GameController) {
	game, controller := setupEmptyGame()
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(1, 3, Neutrino)
	game.SetLocation(3, 1, Neutrino)
	game.SetLocation(3, 3, Neutrino)
	game.State = Player1NeutrinoMove;
	return game, controller
}

/**
 * Setup a game with neutrino in the
 * middle as the only piece and
 * it is player ones turn to move the
 * neutrino
 */
func setupCenteredGame() (*Game, *GameController) {
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1NeutrinoMove;
	return game, controller
}
