/*
给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
注意，删除一个元素后，子数组 不能为空。
n < 1e5   val in [-1e4, 1e4]
*/
/*
最大子数组和， dp[i] = 以nums[i] 为结尾的最大子数组和
dp[i] = max(dp[i-1] + nums[i], nums[i])
删除操作可能发生在哪里？
- 还没有删除过元素。
- 已经删除过一个元素。
定义两个 DP 状态
dp0[i]：以 nums[i] 结尾，且没有执行删除操作的最大子数组和。
dp1[i]：以 nums[i] 结尾，且已经执行过一次删除操作的最大子数组和。
当前元素x = nums[i]
 没有删除元素 dp0[i] = max(dp0[i-1] + x, x)
 已经删除一个
	之前删除  dp1[i-1] + x
	删除当前  dp0[i-1]
dp1[i] = max(dp0[i-1], dp1[i-1] + x)
删除可选  ans = max(dp0[i], dp1[i])
*/
func maximumSum(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// dp[i][0]：以 arr[i] 结尾，没有删除元素的最大子数组和
	// dp[i][1]：以 arr[i] 结尾，已经删除一个元素的最大子数组和
	dp[0][0] = nums[0]
	dp[0][1] = 0
	ans := nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		dp[i][0] = max(dp[i-1][0]+x, x)
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+x)
		ans = max(ans, max(dp[i][0], dp[i][1]))
	}
	return ans
}