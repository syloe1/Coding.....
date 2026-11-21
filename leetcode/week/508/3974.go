/*
给你一个整数数组 nums，以及两个整数 k 和 mul。
从 nums 中选出 恰好 k 个元素。你可以按照任意顺序逐个处理这些元素。
对于每个被选择的元素，都可以 独立地 选择以下两种操作之一：
将该元素的值 加 到总和中；或
将该元素乘以 mul 的 当前 值，并将结果 加 到总和中。
每处理一个被选择的元素后，无论选择哪种操作，mul 都会 减少 1。mul 的当前值可能变为 0 或负数。
返回一个整数，表示可能得到的 最大 总和。
n <= 1e5 v <= 1e5 k < n mul < 1e5

*/

/*
对于每个元素，我们有两种选择：
- 加法：贡献 x
- 乘法：贡献 x * 当前乘数
m是当前乘数
k个乘数 max(1,m),max(1,m-1),…,max(1,m-k+1) 选择最大的k个数


nums 从大到小排序。
选择前 k 个元素。
依次乘以 max(1, mul-i)。
累加答案。
*/
func maxSum(nums []int, k int, mul int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	ans := 0
	for i := 0; i < k; i++ {
		f := max(1, mul-i)
		ans += nums[i] * f
	}
	return int64(ans)
}