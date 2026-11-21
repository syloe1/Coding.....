/*
给你一个整数数组 nums 和一个正整数 k。
你必须选择 nums 的一个 子数组 并执行以下操作之一：
将所选子数组中的每个数字乘以 k。
将所选子数组中的每个数字除以 k。
当正数除以 k 时，除法结果 向下取整。
当负数除以 k 时，除法结果 向上取整。
返回结果数组中 非空子数组的 最大 可能和。
注意，用于执行操作的 子数组 与用于求和的子数组 可以是 不同 的。
子数组 是数组中一段连续的 非空元素序列。
n < 1e5   val in [-1e5,1e5]
k in [1,1e5]
*/
/*
在一个求和子数组中，选择一个连续区间执行操作：
- 不操作
- 乘以 k
- 除以 k
最后求最大子数组和。

注意：操作区间可以为空，即不操作。
*/
/*
我们只需要描述三个阶段：
1. 尚未开始修改。
2. 正在修改。
3. 已经结束修改。
dp[i][0] // 以 nums[i] 结尾，尚未修改
dp[i][1] // 以 nums[i] 结尾，正在修改 nums[i]
dp[i][2] // 以 nums[i] 结尾，之前修改过，当前不修改
- 正数乘以 k，结果更大。
- 负数除以 k，结果更大（向零取整）。
- 零不管怎么操作都不变。
所以我们只需要考虑这一种最优修改方式。

state 0: 未修改
	dp[i][0] = max(dp[i-1][0],0)+x
state 1: 正在修改
从未修改的状态过来 从正在修改的状态过来 直接从当前元素开始
	dp[i][1] = max(dp[i-1][0], dp[i-1][1],0)+y
state 2:结束修改
刚刚结束修改  之前已经结束修改
	dp[i][2] = max(dp[i-1][1], dp[i-1][2]) + x
*/
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	const INF int64 = 1 << 60

	// mul[i][0/1/2]：乘法操作的三个阶段
	// div[i][0/1/2]：除法操作的三个阶段
	mul := make([][3]int64, n)
	div := make([][3]int64, n)

	ans := -INF

	for i := 0; i < n; i++ {
		x := int64(nums[i])
		m := x * int64(k)
		d := x / int64(k)

		if i == 0 {
			mul[i] = [3]int64{x, m, -INF}
			div[i] = [3]int64{x, d, -INF}
		} else {
			mul[i][0] = max(mul[i-1][0], int64(0)) + x
			mul[i][1] = max(max(mul[i-1][0], mul[i-1][1]), int64(0)) + m
			mul[i][2] = max(mul[i-1][1], mul[i-1][2]) + x

			div[i][0] = max(div[i-1][0], int64(0)) + x
			div[i][1] = max(max(div[i-1][0], div[i-1][1]), int64(0)) + d
			div[i][2] = max(div[i-1][1], div[i-1][2]) + x
		}

		ans = max(ans, max(
			max(mul[i][1], mul[i][2]),
			max(div[i][1], div[i][2]),
		))
	}

	return ans
}