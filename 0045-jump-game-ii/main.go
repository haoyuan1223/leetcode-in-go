package _045_jump_game_ii

func jump(nums []int) int {
	maxPos := 0
	end := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
