/*
例 13-1 有序序列中查找元素 (洛谷 P2249)。输入 \(n(n\le 10^6)\) 个不超过 \(10^9\) 的单调不减的 (就是后面的数字不小于前面的数字) 非负整数 \(a_1,a_2,\dots,a_n\)，然后进行 \(m(m\le 10^5)\) 次询问。对于每次询问，给出一个整数 \(q(q\le 10^9)\)，要求输出这个数字在序列中的编号，如果没有找到，则输出 “-1”。
*/
package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	find := func(x int) int {
		l, r := 0, len(a)
		for l < r {
			mid := l + (r-l)/2
			if a[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	for ; t > 0; t-- {
		var x int
		fmt.Scan(&x)
		pos := find(x)
		var res int
		if pos < n && a[pos] == x {
			res = pos + 1
		} else {
			res = -1
		}
		fmt.Printf("%d ", res) // 空格隔开，不换行
	}
}
