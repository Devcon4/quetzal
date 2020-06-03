package main

import "testing"

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestGame_roll_happy_path(t *testing.T) {

	testGame := &Game{}

	type args struct {
		pins int
	}
	tests := []struct {
		name string
		g    Game
		args args
	}{
		{name: "happyPath", g: *testGame, args: args{pins: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.roll(tt.args.pins)
		})
	}
}
func TestGame_roll_panic(t *testing.T) {

	completedGame := &Game{}

	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)
	completedGame.roll(5)

	type args struct {
		pins int
	}
	tests := []struct {
		name string
		g    Game
		args args
	}{
		{name: "too many pins", g: Game{}, args: args{pins: 11}},
		{name: "too many frames", g: *completedGame, args: args{pins: 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shouldPanic(t, func() {
				tt.g.roll(tt.args.pins)
			})
		})
	}
}

func TestGame_score(t *testing.T) {
	tests := []struct {
		name string
		g    Game
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.score(); got != tt.want {
				t.Errorf("Game.score() = %v, want %v", got, tt.want)
			}
		})
	}
}
