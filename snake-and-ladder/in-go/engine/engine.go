package engine

import (
	"fmt"
	"snake-and-ladder/in-go/model"
)

func (engine *engine) AddPlayer(name string) {
	p := model.InitPlayer(name)
	engine.players = append(engine.players, p)
}

func (engine *engine) Play() {
	if len(engine.players) == 0 {
		fmt.Printf("no players joined.")
		return
	}
	totalNoOfMoves := 0
	for {
		totalNoOfMoves++

		currentPlayer := engine.players[0]
		roll := engine.dice.Roll()

		newPosition := currentPlayer.Position + roll
		if newPosition > engine.board.GetEndValue() {
			fmt.Printf("Player: %s, diceRoll: %d\n", currentPlayer, roll)
			engine.players = append(engine.players[1:], engine.players[0])
			continue
		}

		currentPlayer.Position = engine.getNewPosition(currentPlayer.Position + roll)
		fmt.Printf("Player: %s, diceRoll: %d\n", currentPlayer, roll)
		if currentPlayer.Position == engine.board.GetEndValue() {
			currentPlayer.Won = true
			fmt.Printf("%s Won!!!\n", currentPlayer)
			engine.players = engine.players[1:]
		} else {
			engine.players = append(engine.players[1:], engine.players[0])
		}

		if len(engine.players) < 2 {
			fmt.Printf("Game finished!! total moves: %d\n", totalNoOfMoves)
			return
		}
	}
}

func (engine *engine) getNewPosition(pos int64) int64 {
	for _, val := range engine.ladders {
		if val.GetStart() == pos {
			fmt.Printf("is ladder. start: %d, end: %d\n", pos, val)
			return val.GetEnd()
		}
	}
	for _, val := range engine.snakes {
		if val.GetHead() == pos {
			fmt.Printf("is snake. start: %d, end: %d\n", pos, val)
			return val.GetTail()
		}
	}

	return pos
}
