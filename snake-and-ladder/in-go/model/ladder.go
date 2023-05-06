package model

type Ladder struct {
	start int64
	end   int64
}

func InitLadder(start, end int64) *Ladder {
	return &Ladder{
		start: start,
		end:   end,
	}
}

func (ladder *Ladder) GetStart() int64 {
	return ladder.start
}

func (ladder *Ladder) GetEnd() int64 {
	return ladder.end
}
