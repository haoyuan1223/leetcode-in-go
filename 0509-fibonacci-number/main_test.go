package _509_fibonacci_number

import (
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{N: 2},
			want: 1,
		},
		{
			args: args{N: 3},
			want: 2,
		},
		{
			args: args{N: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
