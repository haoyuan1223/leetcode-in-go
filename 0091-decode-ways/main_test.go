package _091_decode_ways

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "12"}, want: 2},
		{args: args{s: "226"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
			if got := numDecodings1(tt.args.s); got != tt.want {
				t.Errorf("numDecodings1() = %v, want %v", got, tt.want)
			}
		})
	}
}
