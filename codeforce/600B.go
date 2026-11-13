/*
对数组 b 中的每个数，统计数组 a 中有多少个元素不超过它，并按顺序输出这些统计结果。

常见做法是：先把数组 a 排序，然后对每个 b[j] 用二分查找 <= b[j] 的元素个数。
我们要在 a 里面定位一个分界点：哪些 a 元素满足条件，哪些不满足。

如果 a 有序，那么所有 <= b[j] 的元素一定集中在左边连续一段，所有 > b[j] 的元素集中在右边连续一段。
这样就能用二分找到“分界点”，分界点左边有多少个元素，答案就是多少。
“找第一个 > x 的位置”的写法（x = b[j]，0-indexed 数组 a）：
目标：找到第一个满足 a[pos] > x 的位置 pos
答案 = pos（因为 pos 前面恰好有 pos 个元素 <= x）
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	//返回第一个 > x 个位置
	upperBound := func(x int) int {
		l, r := 0, n
		for l < r {
			mid := l + (r-l)/2
			if a[mid] <= x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	for j := 0; j < m; j++ {
		var bj int
		fmt.Scan(&bj)
		if j > 0 {
			fmt.Print(" ")
		}
		fmt.Print(upperBound(bj))
	}
	fmt.Println()
}
