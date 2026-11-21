/*
给你一个整数数组 nums 和一个整数数字 x。
如果一个 子数组 nums[l..r] 的元素和同时满足以下两个条件，则认为该子数组是 有效子数组：
该和的首位数字等于 x。
该和的末位数字等于 x。
返回有效子数组的数量。
子数组 是数组中一个连续、非空 的元素序列。
n <= 1500, val < 1e9, x in [1, 9]
*/
/*
子数组和可能非常大，不能直接把和转换成字符串来判断首位和末位。
枚举左右边界 + 前缀和
我们可以枚举每个子数组的左边界 l 和右边界 r。
为了快速计算子数组和，使用前缀和：
sum(l, r) = pre[r + 1] - pre[l]
pre[i] = nums[0] + ... + nums[i-1]
pre[l] = nums[0] +.... + nums[l-1]
pre[r + 1] = nums[0] + ... + nums[r]
*/
func countValidSubarrays(nums []int, x int) int {
	n := len(nums)
	pre := make([]int64, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + int64(nums[i])
	}
	ans := 0
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			s := pre[r+1] - pre[l]
			// 判断末尾
			if s%10 != int64(x) {
				continue
			}
			f := s
			for f >= 10 {
				f /= 10
			}
			if f == int64(x) {
				ans++
			}
		}
	}
	return ans
}