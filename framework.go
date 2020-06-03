package main

type Game struct {
	Frames []Frame
	Points []int
}

type Frame struct {
	Score1 int
	Score2 int
	Score3 int

	IsStrike bool
	IsSpare  bool
}

func (g Game) roll(pins int) {

	if pins > 10 {
		panic("too many pins!")
	}

	if len(g.Frames) > 10 {
		panic("too many frames!")
	}

	g.Points = append(g.Points, pins)
}

func (g Game) score() int {
	return 100
}
