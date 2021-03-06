package _174_dungeon_game

import (
	"testing"
)

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{dungeon: [][]int{
				{1, -2, 3},
				{2, -2, -2},
			}},
			want: 2,
		},
		{
			args: args{dungeon: [][]int{
				{-3, 5},
			}},
			want: 4,
		},
		{
			args: args{dungeon: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
