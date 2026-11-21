/*
在蓝桥王国，数字的大小不仅仅取决于它们的数值大小，还取决于它们所形成的“封闭图形”的个数。

封闭图形是指数字中完全封闭的空间，例如数字 1、2、3、5、7 都没有形成封闭图形，而数字 0、4、6、9 分别形成了 1 个封闭图形，数字 8 则形成了 2个封闭图形。值得注意的是，封闭图形的个数是可以累加的。例如，对于数字68，由于 6 形成了 1 个封闭图形，而 8 形成了 2 个，所以 68 形成的封闭图形的个数总共为 3。
在比较两个数的大小时，如果它们的封闭图形个数不同，那么封闭图形个数较多的数更大。例如，数字 41 和数字 18，它们对应的封闭图形的个数分别为 1 和 2，因此数字 41 小于数组 18。如果两个数的封闭图形个数相同，那么数值较大的数更大。例如，数字 14 和数字 41，它们的封闭图形的个数都是 1，但 14 < 41，所以数字 14 小于数字 41。如果两个数字的封闭图形个数和数值都相同，那么这两个数字被认为是相等的。

小蓝对蓝桥王国的数字大小规则十分感兴趣。现在，他将给定你 n 个数a1, a2, . . . , an，请你按照蓝桥王国的数字大小规则，将这 n 数从小到大排序，并输出排序后结果。
输入的第一行包含一个整数 n ，表示给定的数字个数。

第二行包含 n 个整数 a1, a2, . . . , an ，相邻整数之间使用一个空格分隔，表示待排序的数字。

输出格式
输出一行包含 n 个整数，相邻整数之间使用一个空格分隔，表示按照蓝桥王国的数字大小规则从小到大排序后的结果。
*/
// 预处理  + sort , nlogn
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var ring = [10]int{1, 0, 0, 0, 1, 0, 1, 0, 2, 1}

func calRing(x int) int {
	if x == 0 {
		return ring[0]
	}
	res := 0
	for x > 0 {
		res += ring[x%10]
		x /= 10
	}
	return res
}

type Node struct {
	val int
	r   int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]Node, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].val)
		a[i].r = calRing(a[i].val)
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].r != a[j].r {
			return a[i].r < a[j].r
		}
		return a[i].val < a[j].val
	})

	for i := 0; i < n; i++ {
		if i > 0 {
			out.WriteByte(' ')
		}
		fmt.Fprint(out, a[i].val)
	}
	out.WriteByte('\n')
}
