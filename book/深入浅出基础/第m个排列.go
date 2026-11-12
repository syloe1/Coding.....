/*
1. 从后往前找第一个下标 `i`，满足 `a[i] < a[i+1]`
2. 找不到 → 已经是最后一个排列（本题保证输入合法，不会遇到）
3. 再从后往前找第一个下标 `j`，满足 `a[j] > a[i]`
4. 交换 `a[i]` 和 `a[j]`，反转 i 后面的所有元素
*/
/*
1 2 3 4 5
1 2 3 5 4
1 2 4 3 5
1 2 4 5 3
1 2 5 3 4
1 2 5 4 3   ← 当前
1 3 2 4 5   ← 【真正的下一个】
1 3 2 5 4
1 3 4 2 5
……
1 4 3 2 5
*/
package main

import "fmt"

func nextPermutation(a []int) bool {
	n := len(a)
	// 1. 从后向前找第一个 i，满足 a[i] < a[i+1]
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false //最后一个排列
	}
	// 2. 从后向前找第一个 j，满足 a[j] > a[i]
	j := n - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	//reverse
	l, r := i+1, n-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return true
}
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// 执行 m 次下一个排列
	for t := 0; t < m; t++ {
		nextPermutation(arr)
	}

	// 输出
	for idx := 0; idx < n; idx++ {
		if idx > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[idx])
	}
	fmt.Println()
}
