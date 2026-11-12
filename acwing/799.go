/*
题目：最长无重复元素连续子数组（双指针 / 滑动窗口）
给定一个长度为 n 的整数序列，请找出最长的不包含重复的数的连续区间，输出它的长度。

输入格式：
第一行包含整数 n。
第二行包含 n 个整数（均在 0 ∼ 1e5 范围内），表示整数序列。

输出格式：
共一行，包含一个整数，表示最长的不包含重复的数的连续区间的长度。

数据范围：
1 ≤ n ≤ 100000
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}
	const maxVal = 100000
	cnt := make([]int, maxVal+1)
	l := 0
	ans := 0
	for r := 0; r < n; r++ {
		x := a[r]
		cnt[x]++
		for cnt[x] > 1 {
			cnt[a[l]]--
			l++
		}
		// l r 无重复
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	println(ans)
}
