package _003_Longest_Substring_Without_Repeating_Characters

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "abcabcbb"}, want: 3},
		{args: args{s: "bbbbb"}, want: 1},
		{args: args{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
			if got := lengthOfLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring1() = %v, want %v", got, tt.want)
			}
		})
	}
}
