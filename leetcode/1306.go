/*
给定一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。

请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。

注意，不管是什么情况下，你都无法跳到数组之外。
*/
func canReach(arr []int, start int) bool {
	n := len(arr)
	vs := make([]bool, n)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i < 0 || i >= n || vs[i] {
			return false
		}
		if arr[i] == 0 {
			return true
		}
		vs[i] = true
		return dfs(i+arr[i]) || dfs(i-arr[i])
	}
	return dfs(start)
}
