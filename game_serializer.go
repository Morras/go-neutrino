package neutrino

import (
	"fmt"
	"strconv"
)

var (
	serializerPrefixLength = 11
)

func GameToUInt64(game *Game) uint64 {
	bits := ""

	//The entries are stored by rows, so
	//we must pass through each column before
	//moving on to the next row
	for y := byte(0); y < 5; y++ {
		for x := byte(0); x < 5; x++ {
			entry, _ := game.GetLocation(x, y)
			entryAsBits := strconv.FormatUint(uint64(entry), 2)
			bits += fmt.Sprintf("%02s", entryAsBits)
		}
	}

	stateAsBits := strconv.FormatUint(uint64(game.State), 2)
	bits += fmt.Sprintf("%03s", stateAsBits)
	bits = fmt.Sprintf("%064s", bits)

	output, _ := strconv.ParseUint(bits, 2, 64)
	//todo do something with the error.
	return output
}

func UInt64ToGame(input uint64) *Game {
	game := &Game{}

	bits := strconv.FormatUint(input, 2)
	bits = fmt.Sprintf("%064s", bits)

	var x byte = 0
	var y byte = 0
	for i := 0; i < 50; i += 2 {
		entryAsBits := bits[serializerPrefixLength+i : serializerPrefixLength+i+2]
		//TODO do something with the error
		entry, _ := strconv.ParseUint(entryAsBits, 2, 8)
		game.SetLocation(x, y, Entry(entry))

		x++
		if x > 4 {
			x = 0
			y++
		}

	}

	stateAsBits := bits[61:64]
	//TODO do something with the error
	state, _ := strconv.ParseUint(stateAsBits, 2, 8)
	game.State = State(state)
	return game
}
