package test1

import (
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
