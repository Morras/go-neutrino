package neutrino

/**
* Basic setup of an empty game
* Remember to add a neutrino to your game
* otherwise it is invalid and the game controller
* will panic
 */
func SetupEmptyGame() (*Game, *GameController) {
	game := NewEmptyGame()
	controller := &GameController{}
	controller.StartGame(game)
	return game, controller
}

/**
 * Setup with a player 1 piece on
 * (1,1), (3,1), (3,1) and (3,3)
 * and it is player ones turn to move
 * the neutrino.
 */
func SetupSquaredGame() (*Game, *GameController) {
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
func SetupCenteredGame() (*Game, *GameController) {
	game, controller := SetupEmptyGame()
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1NeutrinoMove
	return game, controller
}
