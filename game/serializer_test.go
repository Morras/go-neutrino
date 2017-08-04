package game

import "testing"

func TestSerialization_SquaredGame(t *testing.T) {
	referenceGame, _ := SetupSquaredGame()

	testSerializationOfGame(referenceGame, t)
}

func TestSerialization_CenteredGame(t *testing.T) {
	referenceGame, _ := SetupCenteredGame()

	testSerializationOfGame(referenceGame, t)
}

func TestSerialization_StandardGame(t *testing.T) {
	referenceGame := NewStandardGame()
	testSerializationOfGame(referenceGame, t)
}

func TestSerialization_RealGame(t *testing.T) {
	referenceGame := NewStandardGame()
	controller := &Controller{}
	controller.NewController(referenceGame)

	testSerializationOfGame(referenceGame, t)
	//Player1
	makeMoveAndCheckError(2, 2, 3, 3, controller, t)
	makeMoveAndCheckError(3, 0, 3, 2, controller, t)
	testSerializationOfGame(referenceGame, t)
	//Player2
	makeMoveAndCheckError(3, 3, 4, 3, controller, t)
	makeMoveAndCheckError(2, 4, 4, 2, controller, t)
	testSerializationOfGame(referenceGame, t)
	//Player1
	makeMoveAndCheckError(4, 3, 0, 3, controller, t)
	makeMoveAndCheckError(2, 0, 3, 0, controller, t)
	testSerializationOfGame(referenceGame, t)
	//Player2
	makeMoveAndCheckError(0, 3, 4, 3, controller, t)
	makeMoveAndCheckError(1, 4, 2, 3, controller, t)
	testSerializationOfGame(referenceGame, t)
	//Player1
	makeMoveAndCheckError(4, 3, 3, 3, controller, t)
	makeMoveAndCheckError(0, 0, 2, 2, controller, t)
	testSerializationOfGame(referenceGame, t)
	//Player2 WIN
	makeMoveAndCheckError(3, 3, 4, 3, controller, t)
	makeMoveAndCheckError(2, 3, 3, 3, controller, t)
	testSerializationOfGame(referenceGame, t)
	if referenceGame.State != Player2Win {
		t.Error("Player 2 should have won by now. Expected", Player2Win, "got", referenceGame.State)
	}
}

func testSerializationOfGame(referenceGame *Game, t *testing.T) {
	intRepresentation := GameToUInt64(referenceGame)
	serializedGame := UInt64ToGame(intRepresentation)
	if result, message := Compare(referenceGame, serializedGame); !result {
		t.Error("Serialized game does not match game that was converted into int: ", message)
	}
}

func makeMoveAndCheckError(fromX, fromY, toX, toY byte, controller *Controller, t *testing.T) {
	_, err := controller.MakeMove(NewMove(fromX, fromY, toX, toY))
	if err != nil {
		t.Error("Invalid move: ", err.Error())
	}
}
