package model

type Snake struct {
	head int64
	tail int64
}

func InitSnake(head, tail int64) *Snake {
	return &Snake{
		head: head,
		tail: tail,
	}
}

func (snake *Snake) GetHead() int64 {
	return snake.head
}

func (snake *Snake) GetTail() int64 {
	return snake.tail
}
