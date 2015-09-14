package neutrino

import "testing"

/** 
 * There is a problem that GameController sends messages
 * to a state and a move channel when a move is being 
 * made, and if these channels are not cleared the 
 * program deadlocks. 
 * This method makes sure to clear the channel and then
 * does nothing
 */
func pollChannels(moveChan <- chan Move, stateChan <- chan State) {
	go pollMoves(moveChan)
	go pollStates(stateChan)
}
func pollMoves(moveChan <- chan Move){
	for move := range moveChan{ move = move}
}

func pollStates(stateChan <- chan State){
	for state := range stateChan{ state = state}
}

/**
	* Basic setup of an empty game
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
 * their own piece.
 */
func setupSquaredGame() (*Game, *GameController) {
	game, controller := setupEmptyGame()
	game.SetLocation(1, 1, Player1)
	game.SetLocation(1, 3, Player1)
	game.SetLocation(3, 1, Player1)
	game.SetLocation(3, 3, Player1)
	game.State = Player1Move;
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

/**
 * Series of test to see if basic movement
 * is working.
 */
func TestMoveNorth(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 2, 0))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	north, err := game.GetLocation(2, 0)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if north != Neutrino {
		t.Error("Expected", Neutrino, "got", north)
	}
}

func TestMoveNorthEast(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 0))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	ne, err := game.GetLocation(4, 0)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if ne != Neutrino {
		t.Error("Expected", Neutrino, "got", ne)
	}
}

func TestMoveEast(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 2))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	east, err := game.GetLocation(4, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if east != Neutrino {
		t.Error("Expected", Neutrino, "got", east)
	}
}

func TestMoveSouthEast(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 4))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	se, err := game.GetLocation(4, 4)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if se != Neutrino {
		t.Error("Expected", Neutrino, "got", se)
	}
}

func TestMoveSouth(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 2, 4))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	sourth, err := game.GetLocation(2, 4)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if sourth != Neutrino {
		t.Error("Expected", Neutrino, "got", sourth)
	}
}

func TestMoveSouthWest(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 4))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	sw, err := game.GetLocation(0, 4)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if sw != Neutrino {
		t.Error("Expected", Neutrino, "got", sw)
	}
}

func TestMoveWest(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 2))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	west, err := game.GetLocation(0, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if west != Neutrino {
		t.Error("Expected", Neutrino, "got", west)
	}
}

func TestMoveNorthWest(t *testing.T){
	game, controller := setupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 0))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	nw, err := game.GetLocation(0, 0)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if nw != Neutrino {
		t.Error("Expected", Neutrino, "got", nw)
	}
}

/**
 * Series of tests to make sure a piece
 * cannot jump over another piece
 */

func TestStopOnPieceNW(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(3, 3, 0, 0))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(3, 3, 2, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceN(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(3, 3, 3, 0))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(3, 3, 3, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceNE(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(1, 3, 4, 0))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(1, 3, 2, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceE(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(1, 1, 4, 1))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(1, 1, 2, 1))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceSE(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(1, 1, 4, 4))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(1, 1, 2, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceS(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(1, 1, 1, 4))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(1, 1, 1, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceSW(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(3, 1, 0, 4))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(3, 1, 2, 2))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

func TestStopOnPieceW(t *testing.T) {
	_, controller := setupSquaredGame()
	//Make invalid move past a piece
	_, moveError := controller.MakeMove(NewMove(3, 3, 0, 3))
	if moveError == nil {
		t.Error("Expected move error for moving past a piece")
	}
	//Make a move that stops on contact with a piece
	_, moveError = controller.MakeMove(NewMove(3, 3, 2, 3))
	if moveError != nil {
		t.Error("Expected no error got", moveError)
	}
}

/**
 * Series of test to make sure you cannot
 * move outside the game board
 * Notice that it is not possible to move to far
 * to the east or to the north as the coordinates
 * is represented in bytes and must be non negative
 */

func TestCannotMoveOutsideBoardW(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 5, 2))
	if moveError == nil {
		t.Error("Expected an error when moving outside the board")
	}
}

func TestCannotMoveOutsideBoardSW(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 5, 5))
	if moveError == nil {
		t.Error("Expected an error when moving outside the board")
	}
}

func TestCannotMoveOutsideBoardS(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 2, 5))
	if moveError == nil {
		t.Error("Expected an error when moving outside the board")
	}
}

/**
 * Series of test to make sure that a piece does
 * not stop in the middle of a path but moves untill
 * it hits something
 */

func TestCannotStopBeforeObstacleN(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 2, 1))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleNE(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 3, 1))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleE(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 3, 2))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleSE(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 3, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleS(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 2, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleSW(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 1, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleW(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 1, 2))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleNW(t *testing.T) {
	_, controller := setupCenteredGame()
	_, moveError := controller.MakeMove(NewMove(2, 2, 1, 1))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

/**
 * A series of tests that makes sure you cannot
 * move a piece if the game is not in the correct
 * state
 */

func TestPlayerOnePieceMustMatchState(t *testing.T) {
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Player1)
	invalidStates := [5]State{Player1NeutrinoMove,
														Player2NeutrinoMove,
														Player2Move,
														Player1Win,
														Player2Win}

	for _, state := range invalidStates {
		game.State = state
		_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
		if moveError == nil {
			t.Error("It should not be possible to move a player1 piece when in state", state)
		}
	}
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
	if moveError != nil {
		t.Error("It should be possible to move player1")
	}
}

func TestPlayerTwoPieceMustMatchState(t *testing.T) {
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Player2)
	invalidStates := [5]State{Player1NeutrinoMove,
														Player2NeutrinoMove,
														Player1Move,
														Player1Win,
														Player2Win}

	for _, state := range invalidStates {
		game.State = state
		_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
		if moveError == nil {
			t.Error("It should not be possible to move a player1 piece when in state", state)
		}
	}
	game.State = Player2Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
	if moveError != nil {
		t.Error("It should be possible to move player2")
	}
}

func TestNeutrinoPieceMustMatchState(t *testing.T) {
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)
	invalidStates := [4]State{Player1Move,
														Player2Move,
														Player1Win,
														Player2Win}

	for _, state := range invalidStates {
		game.State = state
		_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
		if moveError == nil {
			t.Error("It should not be possible to move a player1 piece when in state", state)
		}
	}
	game.State = Player1NeutrinoMove
	_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
	if moveError != nil {
		t.Error("It should be possible to move neutrino in state", Player1NeutrinoMove)
	}
	game.State = Player2NeutrinoMove
	_, moveError = controller.MakeMove(NewMove(0, 0, 0, 4))
	if moveError != nil {
		t.Error("It should be possible to move neutrino in state", Player2NeutrinoMove)
	}
}

func TestEmptyPiecesMustNotBeMoved(t *testing.T) {
	game, controller := setupEmptyGame()
	invalidStates := [6]State{Player1NeutrinoMove,
														Player2NeutrinoMove,
														Player1Move,
														Player2Move,
														Player1Win,
														Player2Win}

	for _, state := range invalidStates {
		game.State = state
		_, moveError := controller.MakeMove(NewMove(2, 2, 0, 0))
		if moveError == nil {
			t.Error("It should not be possible to move from an empty square", state)
		}
	}
}
