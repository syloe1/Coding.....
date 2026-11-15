/*
给你一个整数数组 nums。
如果一个子数组中所有元素的 按位与 运算结果等于该子数组的 最小值 ，则称该子数组为 一致子数组 。
返回 nums 中 一致子数组 的个数。
子数组 是数组中一个连续、非空的元素序列。
n < 100   nums[i] < 100
*/
//固定左端点，向右枚举右端点，同时维护按位与结果和最小值。
func countConsistent(nums []int) int {
	n := len(nums)
	ans := 0
	for l := 0; i < n; i++ {
		mi := nums[l]
		v := nums[l]
		ans++ //长度为1
		for r := l + 1; r < n; r++ {
			v &= nums[r]
			mi = min(mi, nums[r])
			if v == mi {
				ans++
			}
		}
	}
	return ans
}