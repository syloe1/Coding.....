/*
给你一个整数数组 arr 和两个整数 k 和 threshold 。

请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。
*/
func numOfSubarrays(arr []int, k int, threshold int) (cnt int) {
	s := 0
	for i, c := range arr {
		s += c
		if i < k-1 {
			continue
		}
		if s >= k*threshold {
			cnt++
		}
		s -= arr[i-k+1]
	}
	return
}