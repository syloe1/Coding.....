/*
给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最小操作数。

在一次操作中，你可以使数组中的一个元素加 1 或者减 1 。
*/
// 绝对值和最小值 找中位数
func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	var res int64 = 0
	for _, v := range nums {
		diff := v - mid
		if diff < 0 {
			diff = -diff
		}
		res += int64(diff)
	}
	return int(res)
}
