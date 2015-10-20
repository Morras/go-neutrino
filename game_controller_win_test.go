package neutrino

import "testing"

func TestTrappedNeutrinoEast(t  *testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()
	game.SetLocation(0, 0, Player1)
	game.SetLocation(1, 0, Player1)
	game.SetLocation(0, 2, Player2)
	game.SetLocation(1, 2, Player2)
	game.SetLocation(3, 1, Neutrino)
	game.SetLocation(4, 1, Player1)
	game.State = Player1NeutrinoMove

	state, err := controller.MakeMove(NewMove(3, 1, 0, 1))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player1Move {
		t.Error("State should have changed, expected", Player1Move, "got", state)
	}

	state, err = controller.MakeMove(NewMove(4, 1, 1, 1))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player1Win {
		t.Error("Player1 should be winner as neutrino cannot move, expected", Player1Win, "got", state)
	}
}

func TestTrappedNeutrinoWest(t  *testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()
	game.SetLocation(3, 0, Player1)
	game.SetLocation(4, 0, Player1)
	game.SetLocation(3, 2, Player2)
	game.SetLocation(4, 2, Player2)
	game.SetLocation(1, 1, Neutrino)
	game.SetLocation(0, 1, Player2)
	game.State = Player2NeutrinoMove

	state, err := controller.MakeMove(NewMove(1, 1, 4, 1))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player2Move {
		t.Error("State should have changed, expected", Player2Move, "got", state)
	}

	state, err = controller.MakeMove(NewMove(0, 1, 3, 1))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player2Win {
		t.Error("Player2 should be winner as neutrino cannot move, expected", Player2Win, "got", state)
	}
}

func TestTrappedNeutrinoMiddle(t  *testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()
	game.SetLocation(0, 1, Player1)
	game.SetLocation(1, 1, Player1)
	game.SetLocation(2, 1, Player1)
	game.SetLocation(0, 2, Player2)
	game.SetLocation(0, 3, Player2)
	game.SetLocation(1, 3, Player2)
	game.SetLocation(2, 3, Player2)
	game.SetLocation(3, 2, Neutrino)
	game.SetLocation(4, 2, Player1)
	game.State = Player1NeutrinoMove

	state, err := controller.MakeMove(NewMove(3, 2, 1, 2))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player1Move {
		t.Error("State should have changed, expected", Player1Move, "got", state)
	}

	state, err = controller.MakeMove(NewMove(4, 2, 2, 2))
	if err != nil {
		t.Error("It should have been possible to make a move but got", err)
	}
	if state != Player1Win {
		t.Error("Player1 should be winner as neutrino cannot move, expected", Player1Win, "got", state)
	}
}

func TestMoveOwnNeutrinoToP1HomeRowLooses(t* testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()

	game.SetLocation(1, 1, Neutrino)
	game.State = Player1NeutrinoMove
	state, err := controller.MakeMove(NewMove(1, 1, 1, 0))
	if err != nil {
		t.Error("Expected to be able to make a move, got", err)
	}
	if state != Player2Win {
		t.Error("Expected", Player2Win, " but state was", state)
	}
}

func TestMoveOwnNeutrinoToP1HomeRowWins(t* testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()

	game.SetLocation(4, 1, Neutrino)
	game.State = Player2NeutrinoMove
	state, err := controller.MakeMove(NewMove(4, 1, 4, 0))
	if err != nil {
		t.Error("Expected to be able to make a move, got", err)
	}
	if state != Player2Win {
		t.Error("Expected", Player2Win, " but state was", state)
	}
}

func TestMoveOwnNeutrinoToP2HomeRowLooses(t* testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()

	game.SetLocation(0, 1, Neutrino)
	game.State = Player2NeutrinoMove
	state, err := controller.MakeMove(NewMove(0, 1, 0, 4))
	if err != nil {
		t.Error("Expected to be able to make a move, got", err)
	}
	if state != Player1Win {
		t.Error("Expected", Player1Win, " but state was", state)
	}
}

func TestMoveOwnNeutrinoToP2HomeRowWins(t* testing.T){
	game, controller := setupEmptyGame()
	defer controller.EndGame()

	game.SetLocation(2, 1, Neutrino)
	game.State = Player1NeutrinoMove
	state, err := controller.MakeMove(NewMove(2, 1, 2, 4))
	if err != nil {
		t.Error("Expected to be able to make a move, got", err)
	}
	if state != Player1Win {
		t.Error("Expected", Player1Win, " but state was", state)
	}
}
