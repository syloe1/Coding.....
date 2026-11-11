/*
给你一个整数数组 nums。

在一步操作中，你可以将任意单个元素 nums[i] 的值 增加 1。

返回使数组中的所有元素都 相等 所需的 最小总操作次数 。
*/
//最优目标值，数组最大值
func minMoves(nums []int) int {
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	var res int64 = 0
	for _, v := range nums {
		res += int64(maxVal - v)
	}
	return int(res)
}