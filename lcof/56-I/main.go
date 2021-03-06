package _6_I

// 除了2个数，其他都出现了2次
// 根据异或算法的特点：相同为0，相异为1
// 所有数的异或结果，即为2个只出现一次的数的异或结果，且不为0
// eg：假设数组异或的二进制结果为10010，那么说明这两个数从右向左数第2位是不同的
// 那么可以根据数组里面所有数的第二位为0或者1将数组划分为2组
// 这样做可以将目标数必然分散在不同的组，而且第2位相同的数必落在同组
// 这两个组里面的数各自进行异或，得到的结果就是答案
func singleNumbers(nums []int) []int {
	a, b, x := 0, 0, 0
	for _, num := range nums {
		x ^= num
	}
	x = x & -x
	for _, num := range nums {
		if num&x != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
