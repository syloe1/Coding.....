/*

伐木工人 Mirko 需要砍 M 米长的木材。
他有一台伐木机：设置一个高度 h，锯片升到 h，每棵树高于 h 的部分会被锯下来收集，不高于 h 的部分保留不动。

现在一共有 n 棵树，高度分别为 a₁,a₂,…,aₙ。
求最大的整数高度 h，使得收集到的木材总长度至少为 m。

限制：
n ≤ 10⁶，每棵树高度不超过 10⁹

输入格式
第一行两个整数 n,m。
第二行 n 个整数，表示每棵树的高度。

输出格式
输出满足条件的最大 h。
*/
// h 越大，锯下来的木材总量越少；h 越小，锯下来的木材越多。
// check函数 sum = Σ max(a[i] - mid, 0)
package main

import "fmt"

func check(a []int64, h int64, m int64) bool {
	var sum int64
	for _, v := range a {
		if v > h {
			sum += v - h
		}
		if sum >= m {
			return true
		}
	}
	return sum >= m
}
func main() {
	var n int
	var m int64
	fmt.Scan(&n, &m)
	a := make([]int64, n)
	var maxH int64
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > maxH {
			maxH = a[i]
		}
	}
	l, r := int64(0), maxH+1
	for l < r {
		mid := l + (r-l)/2
		if check(a, mid, m) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(l - 1)
}
