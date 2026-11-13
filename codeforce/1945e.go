/*
先在原数组上模拟二分，得到最终位置 l0。
如果 p[l0] == x，直接输出 0，结束。

找到 x 当前所在位置 posx，交换 l0 和 posx，把 x 放到 l0。
记录这次操作 (l0, posx)。

用交换后的数组再模拟一次二分，得到新的最终位置 l1。
如果 p[l1] == x，说明一次交换就够了，输出 1 和 (l0, posx)。

如果 p[l1] != x，说明第一次交换改变了二分路径，最终位置跑到了 l1。
此时 x 在 l0，直接交换 l1 和 l0，把 x 移到 l1。
输出 2，操作是 (l0, posx) 和 (l1, l0)。

为什么最多两次一定够
第一次交换把 x 放到原最终位置 l0。如果路径没变，一次就成。如果路径变了，新的最终位置是 l1，第二次交换把 x 从 l0 移到 l1，此时 l1 满足二分结束条件，x 就停在 l1 上。题目保证两次足够。
*/
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, x int
		fmt.Scan(&n, &x)
		p := make([]int, n+1)
		posX := 0
		for i := 1; i <= n; i++ {
			fmt.Scan(&p[i])
			if p[i] == x {
				posX = i
			}
		}
		bsearch := func() int {
			l, r := 1, n+1
			for r-l != 1 {
				mid := l + (r-l)/2
				if p[mid] <= x {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}

		l0 := bsearch()
		if p[l0] == x {
			fmt.Println(0)
			continue
		}
		p[l0], p[posX] = p[posX], p[l0]
		l1 := bsearch()
		if p[l1] == x {
			fmt.Println(1)
			fmt.Println(l0, posX)
			continue
		}
		fmt.Println(2)
		fmt.Println(l0, posX)
		fmt.Println(l1, l0)
	}
}
