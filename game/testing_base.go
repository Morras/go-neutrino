package game

/**
* Basic setup of an empty game
* Remember to add a neutrino to your game
* otherwise it is invalid and the game controller
* will panic
 */
func SetupEmptyGame() (*Game, *Controller) {
	game := NewEmptyGame()
	controller := &Controller{}
	controller.StartGame(game)
	return game, controller
}

/**
 * Setup with a player 1 piece on
 * (1,1), (3,1), (3,1) and (3,3)
 * and it is player ones turn to move
 * the neutrino.
 */
func SetupSquaredGame() (*Game, *Controller) {
	game, controller := SetupEmptyGame()
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(1, 3, Neutrino)
	game.SetLocation(3, 1, Neutrino)
	game.SetLocation(3, 3, Neutrino)
	game.State = Player1NeutrinoMove
	return game, controller
}

/**
 * Setup a game with neutrino in the
 * middle as the only piece and
 * it is player ones turn to move the
 * neutrino
 */
func SetupCenteredGame() (*Game, *Controller) {
	game, controller := SetupEmptyGame()
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1NeutrinoMove
	return game, controller
}
