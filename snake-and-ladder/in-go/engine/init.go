package engine

import (
	"fmt"
	"math/rand"
	"snake-and-ladder/in-go/model"
)

func InitEngine(numberOfSnakes, numberOfLadders, boardSize int64) *engine {
	// map to store start and end positions of snakes and ladders
	// so that they dont overlap
	snakeLadderMap := make(map[string]bool)
	snakes := make([]*model.Snake, 0, numberOfSnakes)
	ladders := make([]*model.Ladder, 0, numberOfLadders)

	for i := 0; i < int(numberOfSnakes); i++ {
		for {
			head := rand.Int63n(boardSize) + 1
			tail := rand.Int63n(boardSize) + 1
			if tail >= head || head == boardSize {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", head, tail)]; !ok {
				snakes = append(snakes, model.InitSnake(head, tail))
				snakeLadderMap[fmt.Sprintf("%d:%d", head, tail)] = true
				break
			}
		}
	}

	for i := 0; i < int(numberOfLadders); i++ {
		for {
			start := rand.Int63n(boardSize) + 1
			end := rand.Int63n(boardSize) + 1
			if start >= end || start == 1 {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
				ladders = append(ladders, model.InitLadder(start, end))
				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}
		}
	}

	return &engine{
		board:   model.InitBoard(boardSize),
		dice:    model.InitDice(6),
		snakes:  snakes,
		ladders: ladders,
		players: []*model.Player{},
	}
}
