/*
给你一个长度为偶数 n 的整数数组 nums。

你需要将 nums 中的元素两两合并：将 nums[0] 与 nums[1] 合并，将 nums[2] 与 nums[3] 合并，依此类推，直到将 nums[n - 2] 与 nums[n - 1] 合并。

返回所有合并后得到的整数中的最小值。

合并两个整数，是指将第二个整数的所有十进制数位拼接在第一个整数的所有十进制数位之后。
2 <= n == nums.length <= 100
n 为偶数
1 <= nums[i] <= 1000
*/
package main

import (
	"math"
	"strconv"
)

func minMergedInteger(nums []int) int {
	mi := math.MaxInt
	for i := 0; i < len(nums); i += 2 {
		a := nums[i]
		b := nums[i+1]
		s := strconv.Itoa(a) + strconv.Itoa(b)
		num, _ := strconv.Atoi(s)
		if num < mi {
			mi = num
		}
	}
	return mi
}
