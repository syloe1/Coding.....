/*
上面的布局可以用序列 2 4 6 1 3 5 来描述，第 i 个数字表示在第 i 行的相应位置有一个棋子，如下：

行号 1 2 3 4 5 6

列号 2 4 6 1 3 5

这只是棋子放置的一个解。请编一个程序找出所有棋子放置的解。
并把它们以上面的序列方法输出，解按字典顺序排列。
请输出前 3 个解。最后一行是解的总个数。

输入格式
一行一个正整数 n，表示棋盘是 n×n 大小的。

输出格式
前三行为前三个解，每个解的两个数字之间用一个空格隔开。第四行只有一个数字，表示解的总数。

输入输出样例
输入 #1复制

6
输出 #1复制

2 4 6 1 3 5
3 6 2 5 1 4
4 1 5 2 6 3
4
说明/提示
【数据范围】
对于 100% 的数据，6≤n≤13。
*/

package main

import "fmt"

var (
	n     int
	total int
	col   []bool // col[c] 第c列是否被占用
	dg    []bool // 主对角线
	udg   []bool // 副对角线
	path  []int  //path[i] 第i行皇后放在第几列
)

// dfs(r) 处理第r行
func dfs(r int) {
	// 所有行放完， 得到一组解
	if r > n {
		total++
		if total <= 3 {
			for i := 1; i <= n; i++ {
				if i > 1 {
					fmt.Print(" ")
				}
				fmt.Print(path[i])
			}
			fmt.Println()
		}
		return
	}
	// 枚举
	for c := 1; c <= n; c++ {
		diag1 := r - c + n
		diag2 := r + c
		// 列、两条对角线都没有皇后，才可以放
		if !col[c] && !dg[diag1] && !udg[diag2] {
			col[c] = true
			dg[diag1] = true
			udg[diag2] = true
			path[r] = c

			dfs(r + 1)

			// 回溯：撤销标记，恢复现场
			col[c] = false
			dg[diag1] = false
			udg[diag2] = false
		}
	}
}
func main() {
	fmt.Scan(&n)
	size := 2 * n
	col = make([]bool, n+1)
	dg = make([]bool, size+2)
	udg = make([]bool, size+2)
	path = make([]int, n+1)

	dfs(1)
	fmt.Println(total)
}
