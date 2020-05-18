package _1

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j]%2 == 0 {
			j--
		}
		for i < j && nums[i]%2 == 1 {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
