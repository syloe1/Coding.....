/*
给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。

*/
func findMaxAverage(nums []int, k int) float64 {
	ans := math.MinInt64
	x := 0
	for i, c := range nums {
		x += c
		if i < k-1 {
			continue
		}
		ans = max(ans, x)
		x -= nums[i-k+1]
	}
	return float64(ans) / float64(k)
}