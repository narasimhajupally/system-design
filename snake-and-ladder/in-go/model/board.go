package model

type Board struct {
	size  int64
	start int64
	end   int64
}

func InitBoard(size int64) *Board {
	return &Board{
		size:  size,
		start: 1,
		end:   size,
	}
}

func (b *Board) GetEndValue() int64 {
	return b.size
}
