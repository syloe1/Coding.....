/*
## 题目描述

C 国由 n 个小岛组成，为了方便小岛之间联络，C 国在小岛间建立了 m 座大桥，每座大桥连接两座小岛。两个小岛间可能存在多座桥连接。然而，由于海水冲刷，有一些大桥面临着不能使用的危险。

如果两个小岛间的所有大桥都不能使用，则这两座小岛就不能直接到达了。但是，只要这两座小岛的居民能通过其他的桥或者其他的小岛互相到达，他们就会安然无事。但是，如果前一天两个小岛之间还有方法可以到达，后一天却不能到达了，居民们就会一起抗议。

现在 C 国的国王已经知道了每座桥能使用的天数，超过这个天数就不能使用了。现在他想知道居民们会有多少天进行抗议。

下文中的样例说明
第一天后 2 和 3 之间的桥不能使用，不影响。
第二天后 1 和 2 之间，以及 1 和 3 之间的桥不能使用，居民们会抗议。
第三天后 3 和 4 之间的桥不能使用，居民们会抗议。

## 输入格式

输入的第一行包含两个整数 n, m，分别表示小岛的个数和桥的数量。
接下来 m 行，每行三个整数 a, b, t，分别表示该座桥连接 a 号和 b 号两个小岛，能使用 t 天。小岛的编号从 1 开始递增。

## 数据规模和约定

对于 100% 的数据， 1 <= n <= 10000， 1 <= m <= 100000， 1 <= a, b <= n， 1 <= t <= 100000。

## 输出格式

输出一个整数，表示居民们会抗议的天数。
*/
/*
从所有桥都已经失效的状态开始，按照时间从大到小恢复桥梁。
1. 所有桥都失效，只有孤立岛屿。
2. 恢复失效时间最长的桥。
3. 再恢复失效时间次长的桥。
4. 继续恢复，直到所有桥都恢复。
恢复桥梁只需要并查集合并，非常适合并查集。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Edge struct {
	a, b int
	t    int
}

var fa []int

func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}
func union(a, b int) bool {
	x := find(a)
	y := find(b)
	if x == y {
		return false
	}
	fa[x] = y
	return true
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	edges := make([]Edge, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &edges[i].a, &edges[i].b, &edges[i].t)
	}
	// 按失效时间从大到小排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].t > edges[j].t
	})
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
	ans := 0
	/*
		第t天桥断开， 多出连通块，居民抗议
		倒着来， t的时候，桥被加回来， 连通块合并
	*/
	// 从大到小恢复桥梁
	for i := 0; i < m; {
		j := i
		changed := false

		// 处理失效时间相同的所有桥
		for j < m && edges[j].t == edges[i].t {
			if union(edges[j].a, edges[j].b) {
				changed = true
			}
			j++
		}

		// 只要这个时间段有桥真正连接了两个连通分量，
		// 就说明对应的失效日发生过抗议
		if changed {
			ans++
		}

		i = j
	}
	fmt.Fprintln(out, ans)
}
