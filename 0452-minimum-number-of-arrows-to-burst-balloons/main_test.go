package _452_minimum_number_of_arrows_to_burst_balloons

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, want: 2},
		{args: args{points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, want: 4},
		{args: args{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}}, want: 2},
		{args: args{points: [][]int{{1, 2}}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
