package engine

import "snake-and-ladder/in-go/model"

type engine struct {
	board   *model.Board
	dice    *model.Dice
	snakes  []*model.Snake
	ladders []*model.Ladder
	players []*model.Player
}
