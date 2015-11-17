package game

import "testing"

/**
 * Series of test to see if basic movement
 * is working.
 */
func TestMoveNorth(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 2, 0))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveNorthEast(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 0))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveEast(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 2))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveSouthEast(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 4, 4))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveSouth(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 2, 4))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveSouthWest(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 4))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveWest(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 2))

	middle, err := game.GetLocation(2, 2)
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

func TestMoveNorthWest(t *testing.T) {
	game, controller := SetupCenteredGame()

	controller.MakeMove(NewMove(2, 2, 0, 0))

	middle, err := game.GetLocation(2, 2)
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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupSquaredGame()

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
	_, controller := SetupCenteredGame()

	_, moveError := controller.MakeMove(NewMove(2, 2, 5, 2))
	if moveError == nil {
		t.Error("Expected an error when moving outside the board")
	}
}

func TestCannotMoveOutsideBoardSW(t *testing.T) {
	_, controller := SetupCenteredGame()

	_, moveError := controller.MakeMove(NewMove(2, 2, 5, 5))
	if moveError == nil {
		t.Error("Expected an error when moving outside the board")
	}
}

func TestCannotMoveOutsideBoardS(t *testing.T) {
	_, controller := SetupCenteredGame()

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
	game, controller := SetupEmptyGame()

	game.SetLocation(2, 3, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 3, 2, 1))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleNE(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(1, 3, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(1, 3, 3, 1))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleE(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(1, 2, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 3, 2))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleSE(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(1, 1, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 3, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleS(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(2, 1, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 2, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleSW(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(3, 1, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 1, 3))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleW(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(3, 2, Player1)
	game.State = Player1Move
	_, moveError := controller.MakeMove(NewMove(2, 2, 1, 2))
	if moveError == nil {
		t.Error("Expected an error when stopping a piece before an obstacle")
	}
}

func TestCannotStopBeforeObstacleNW(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(3, 3, Player1)
	game.State = Player1Move
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
	game, controller := SetupEmptyGame()

	game.SetLocation(2, 2, Player1)
	game.SetLocation(4, 4, Neutrino)
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
	game, controller := SetupEmptyGame()

	game.SetLocation(2, 2, Player2)
	game.SetLocation(4, 4, Neutrino)
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
	game, controller := SetupEmptyGame()

	game.SetLocation(2, 2, Neutrino)
	game.SetLocation(4, 4, Neutrino)
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
	game, controller := SetupEmptyGame()

	game.SetLocation(4, 4, Neutrino)
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

/**
 * Series of test to make sure a piece is only moved
 * in a horizontal, vertical or diagonal direction
 */

func TestCanOnlyMoveStraightLowerLeftStart(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(1, 3, Player1)
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1Move
	invalidMoves := [10]Move{NewMove(1, 3, 0, 0),
		NewMove(1, 3, 0, 1),
		NewMove(1, 3, 2, 0),
		NewMove(1, 3, 2, 1),
		NewMove(1, 3, 3, 0),
		NewMove(1, 3, 3, 2),
		NewMove(1, 3, 3, 4),
		NewMove(1, 3, 4, 1),
		NewMove(1, 3, 4, 2),
		NewMove(1, 3, 4, 4)}

	for _, move := range invalidMoves {
		_, moveError := controller.MakeMove(move)
		if moveError == nil {
			t.Error("A piece should only be able to move in straight lines", move)
		}
	}
}

func TestCanOnlyMoveStraightUpperRightStart(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.SetLocation(3, 1, Player1)
	game.SetLocation(2, 2, Neutrino)
	game.State = Player1Move
	invalidMoves := [10]Move{NewMove(3, 1, 0, 0),
		NewMove(3, 1, 0, 2),
		NewMove(3, 1, 0, 3),
		NewMove(3, 1, 1, 0),
		NewMove(3, 1, 1, 2),
		NewMove(3, 1, 1, 4),
		NewMove(3, 1, 2, 3),
		NewMove(3, 1, 2, 4),
		NewMove(3, 1, 4, 3),
		NewMove(3, 1, 4, 4)}

	for _, move := range invalidMoves {
		_, moveError := controller.MakeMove(move)
		if moveError == nil {
			t.Error("A piece should only be able to move in straight lines", move)
		}
	}
}

/**
 * A piece cannot be moved to its own location
 */

func TestPieceMustMoveToAnotherLocation(t *testing.T) {
	_, controller := SetupCenteredGame()

	_, moveError := controller.MakeMove(NewMove(2, 2, 2, 2))
	if moveError == nil {
		t.Error("It should not be possible to move a piece to its own location")
	}
}

/**
 * Tests to make sure the game state is
 * advanced upon a move
 */

func TestGameStateShouldAdvanceUponMove(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player1NeutrinoMove
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(1, 2, Player1)
	game.SetLocation(1, 3, Player2)

	//Move neutrino piece
	state, moveError := controller.MakeMove(NewMove(1, 1, 0, 1))
	if moveError != nil {
		t.Error("It should be possible to move neutrino piece:", moveError)
	}
	if state != Player1Move {
		t.Error("State should have advanced after neutrino was moved but was. Expected", Player1Move, "but was", state)
	}
	//Move player 1 piece
	state, moveError = controller.MakeMove(NewMove(1, 2, 0, 2))
	if moveError != nil {
		t.Error("It should be possible to move player 1 piece:", moveError)
	}
	if state != Player2NeutrinoMove {
		t.Error("State should have advanced after player 1 piece was moved. Expected", Player2NeutrinoMove, "but was", state)
	}
	//Move neutrino piece
	state, moveError = controller.MakeMove(NewMove(0, 1, 4, 1))
	if moveError != nil {
		t.Error("It should be possible to move neutrino piece:", moveError)
	}
	if state != Player2Move {
		t.Error("State should have advanced after neutrino was moved. Expected", Player2Move, "but was", state)
	}
	//Move player 2 piece
	state, moveError = controller.MakeMove(NewMove(1, 3, 0, 3))
	if moveError != nil {
		t.Error("It should be possible to move player 2 piece:", moveError)
	}
	if state != Player1NeutrinoMove {
		t.Error("State should have advanced after neutrino was moved. Expected", Player1NeutrinoMove, "but was", state)
	}
}

func TestGameStateShouldNotAdvanceUponInvalidMove(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player1NeutrinoMove
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(2, 1, Player1)

	//Try a moving in a non straight line
	controller.MakeMove(NewMove(1, 1, 0, 3))
	if game.State != Player1NeutrinoMove {
		t.Error("State should not change on an illegal move, expected", Player1NeutrinoMove, "was", game.State)
	}
	//Try stopping before the edge
	controller.MakeMove(NewMove(1, 1, 1, 3))
	if game.State != Player1NeutrinoMove {
		t.Error("State should not change on an illegal move, expected", Player1NeutrinoMove, "was", game.State)
	}
	//Try jumping over a piece
	controller.MakeMove(NewMove(1, 1, 4, 1))
	if game.State != Player1NeutrinoMove {
		t.Error("State should not change on an illegal move, expected", Player1NeutrinoMove, "was", game.State)
	}
	//Try moving a wrong piece
	controller.MakeMove(NewMove(2, 1, 2, 0))
	if game.State != Player1NeutrinoMove {
		t.Error("State should not change on an illegal move, expected", Player1NeutrinoMove, "was", game.State)
	}
}

//We cannot make any moves from the west or north of
//the board as the indices are bytes
func TestMoveFromOutsideTheBoardIsInvalid(t *testing.T) {
	//We cannot add any pieces because the game
	//will give an error when we try
	_, controller := SetupEmptyGame()

	_, err := controller.MakeMove(NewMove(1, 6, 1, 0))
	if err == nil {
		t.Error("Expected an error when making a move from south of the board")
	}
	_, err = controller.MakeMove(NewMove(6, 1, 0, 1))
	if err == nil {
		t.Error("Expected an error when making a move from east of the board")
	}
	_, err = controller.MakeMove(NewMove(6, 6, 0, 0))
	if err == nil {
		t.Error("Expected an error when making a move from south-east of the board")
	}
}

func TestCannotMoveToAnotherPiece(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player1NeutrinoMove
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(4, 1, Player1)

	_, err := controller.MakeMove(NewMove(1, 1, 4, 1))
	if err == nil {
		t.Error("Expected an error when making a move from south of the board")
	}
}

func TestCannotMoveAllPiecesBackToPlayer1HomeRow(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player1Move
	game.SetLocation(0, 0, Player1)
	game.SetLocation(1, 0, Player1)
	game.SetLocation(2, 0, Player1)
	game.SetLocation(3, 0, Player1)
	game.SetLocation(4, 0, Player1)
	game.SetLocation(3, 3, Neutrino)

	_, err := controller.MakeMove(NewMove(1, 0, 1, 4))
	if err != nil {
		t.Error("Expected to be able to make a straight move, got", err)
	}
	game.State = Player1Move
	//You cannot have your own five pieces back on your own home
	//row after it has been broken by the first move
	_, err = controller.MakeMove(NewMove(1, 4, 1, 0))
	if err == nil {
		t.Error("Expected to not be able to return all pieces to home row")
	}
	//If we remove a piece from the home work it should then be
	//possible to move back
	game.State = Player1Move
	_, err = controller.MakeMove(NewMove(4, 0, 4, 4))
	if err != nil {
		t.Error("Expected to be able to make a straight move, got", err)
	}
	game.State = Player1Move
	_, err = controller.MakeMove(NewMove(1, 4, 1, 0))
	if err != nil {
		t.Error("Expected to be able move back when there is only three pieces on the home row, got", err)
	}
	//An opposing piece does not count against moving to the home row
	game.SetLocation(4, 0, Player2)
	game.State = Player1Move
	controller.MakeMove(NewMove(1, 0, 1, 4))
	game.State = Player1Move
	_, err = controller.MakeMove(NewMove(1, 4, 1, 0))
	if err != nil {
		t.Error("Expected to be able move back when there is only three pieces and an opposing piece on the home row, got", err)
	}
}

func TestPlayer1CanMoveWithinHomeRowWithFourPieces(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player1Move
	game.SetLocation(0, 0, Player1)
	game.SetLocation(1, 0, Player1)
	game.SetLocation(2, 0, Player1)
	game.SetLocation(4, 0, Player1)
	game.SetLocation(3, 3, Neutrino)

	//With only four pieces on your home row you should be
	//able to move within the home row
	_, err := controller.MakeMove(NewMove(2, 0, 3, 0))
	if err != nil {
		t.Error("Expected to be able to move within own home row when only four pieces are present")
	}
}

func TestCannotMoveAllPiecesBackToPlayer2HomeRow(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player2Move
	game.SetLocation(0, 4, Player2)
	game.SetLocation(1, 4, Player2)
	game.SetLocation(2, 4, Player2)
	game.SetLocation(3, 4, Player2)
	game.SetLocation(4, 4, Player2)
	game.SetLocation(3, 3, Neutrino)

	_, err := controller.MakeMove(NewMove(1, 4, 1, 0))
	if err != nil {
		t.Error("Expected to be able to make a straight move, got", err)
	}
	game.State = Player2Move
	//You cannot have your own five pieces back on your own home
	//row after it has been broken by the first move
	_, err = controller.MakeMove(NewMove(1, 0, 1, 4))
	if err == nil {
		t.Error("Expected to not be able to return all pieces to home row")
	}
	//If we remove a piece from the home work it should then be
	//possible to move back
	game.State = Player2Move
	_, err = controller.MakeMove(NewMove(4, 4, 4, 0))
	if err != nil {
		t.Error("Expected to be able to make a straight move, got", err)
	}
	game.State = Player2Move
	_, err = controller.MakeMove(NewMove(1, 0, 1, 4))
	if err != nil {
		t.Error("Expected to be able move back when there is only three pieces on the home row, got", err)
	}
	//An opposing piece does not count against moving to the home row
	game.SetLocation(4, 4, Player1)
	game.State = Player2Move
	controller.MakeMove(NewMove(1, 4, 1, 0))
	game.State = Player2Move
	_, err = controller.MakeMove(NewMove(1, 0, 1, 4))
	if err != nil {
		t.Error("Expected to be able move back when there is only three pieces and an opposing piece on the home row, got", err)
	}
}

func TestPlayer2CanMoveWithinHomeRowWithFourPieces(t *testing.T) {
	game, controller := SetupEmptyGame()

	game.State = Player2Move
	game.SetLocation(0, 4, Player2)
	game.SetLocation(1, 4, Player2)
	game.SetLocation(2, 4, Player2)
	game.SetLocation(4, 4, Player2)
	game.SetLocation(3, 3, Neutrino)

	//With only four pieces on your home row you should be
	//able to move within the home row
	_, err := controller.MakeMove(NewMove(2, 4, 3, 4))
	if err != nil {
		t.Error("Expected to be able to move within own home row when only four pieces are present")
	}
}
