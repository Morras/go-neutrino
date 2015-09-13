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
 * Series of test to see if basic movement
 * is working.
 */
func TestMoveNorth(t *testing.T){
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

	controller.MakeMove(NewMove(2, 2, 4, 0))

	middle, err  := game.GetLocation(2, 2)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if middle != EmptySquare {
		t.Error("Expected", EmptySquare, "got", middle)
	}

	east, err := game.GetLocation(4, 0)
	if err != nil {
		t.Error("Expected no error got", err)
	}
	if east != Neutrino {
		t.Error("Expected", Neutrino, "got", east)
	}
}

func TestMoveSouthEast(t *testing.T){
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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
	game, controller := setupEmptyGame()
	game.SetLocation(2, 2, Neutrino)

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

func TestFailing(t *testing.T){
	t.Error("Test failed on purpose")
}

