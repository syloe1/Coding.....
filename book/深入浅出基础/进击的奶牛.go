/*
题目描述：
一个牛棚有 n 个隔间，它们分布在一条直线上，坐标是 \(x_1,x_2,\dots,x_n\)。
现在需要把 c 头牛安置到某些隔间，使得所有牛中相邻两头的最近距离越大越好，求这个最大的最近距离。

输入格式：
第一行两个整数 n，c。
第二行 n 个整数，代表各个隔间的坐标。

输出格式：
输出最大的最近距离。

>
> 考点：二分答案（经典最大化最小值）
> 单调性：距离 d 越大，能放下的牛的数量越少；d 越小，能放下的牛越多。
> check (d)：判断隔间里能不能放下至少 c 头牛，满足任意两头牛之间距离≥d。
*/
/*
先想单调性 → 再写 check 函数 → 再确定二分区间、循环不变量、模板分支
- d 越大 → 能放下的牛越少
- d 越小 → 能放下的牛越多
*/
package main

import (
	"fmt"
	"sort"
)

func check(pos []int, d int, c int) bool {
	count := 1
	last := pos[0] //记录上一头牛的位置
	for i := 1; i < len(pos); i++ {
		if pos[i]-last >= d {
			count++ //放一头牛
			last = pos[i]
			if count >= c {
				return true
			}
		}
	}
	return count >= c
}
func main() {
	var n, c int
	fmt.Scan(&n, &c)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pos[i])
	}
	sort.Ints(pos)

	l, r := 1, pos[n-1]-pos[0]+1
	for l < r {
		mid := l + (r-l)/2
		if check(pos, mid, c) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(l - 1)
}
