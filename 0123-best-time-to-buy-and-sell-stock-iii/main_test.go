package _123_best_time_to_buy_and_sell_stock_iii

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}}, want: 6},
		{args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
			if got := maxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit1() = %v, want %v", got, tt.want)
			}
		})
	}
}