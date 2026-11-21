/*
# 题目描述

小红拿到了一个字符矩阵，矩阵中仅包含 r、e、d 这三种字符。

小红每次操作可以将任意字符修改为 r、e、d 这三种字符中的一种。她希望最终任意两个相邻的字母都不相同（相邻指上下左右相邻）。小红想知道，至少需要修改多少个字符？

## 输入描述

第一行输入两个正整数 n, m，代表矩阵的行数和列数。
接下来 n 行，每行输入一个长度为 m 的、仅由 r、e、d 组成的字符串。
约束：
1 ≤ n ≤ 4
1 ≤ m ≤ 1000

## 输出描述

输出一个整数，代表需要修改的字母数量的最小值。
*/
/*
行数很少，列数比较大。这通常暗示我们考虑列状态.
把较短的一维压缩成状态，然后沿着较长的一维做 DP。
dp[j][s] 处理到第j列并且第j列最终变成状态s时， 最少要修改多少个字符
t是上一列状态
dp[j][s] = cost(j, s) + min(dp[j-1][t])
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e9)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &g[i])
	}

	// 字符映射：r=0, e=1, d=2
	mp := map[byte]int{
		'r': 0,
		'e': 1,
		'd': 2,
	}

	// 枚举所有合法列状态
	total := 1
	for i := 0; i < n; i++ {
		total *= 3
	}

	states := make([][]int, 0)

	for s := 0; s < total; s++ {
		x := s
		state := make([]int, n)

		valid := true

		for i := 0; i < n; i++ {
			state[i] = x % 3
			x /= 3

			// 上下相邻不能相同
			if i > 0 && state[i] == state[i-1] {
				valid = false
				break
			}
		}

		if valid {
			states = append(states, state)
		}
	}

	k := len(states)

	// 预处理状态之间是否可以相邻
	compatible := make([][]bool, k)
	for i := 0; i < k; i++ {
		compatible[i] = make([]bool, k)
		for j := 0; j < k; j++ {
			ok := true

			for p := 0; p < n; p++ {
				if states[i][p] == states[j][p] {
					ok = false
					break
				}
			}

			compatible[i][j] = ok
		}
	}

	// 预处理每一列变成每种状态的修改代价
	cost := make([][]int, m)
	for col := 0; col < m; col++ {
		cost[col] = make([]int, k)

		for s := 0; s < k; s++ {
			cnt := 0

			for row := 0; row < n; row++ {
				if mp[g[row][col]] != states[s][row] {
					cnt++
				}
			}

			cost[col][s] = cnt
		}
	}

	// 滚动 DP
	dp := make([]int, k)
	ndp := make([]int, k)

	for s := 0; s < k; s++ {
		dp[s] = cost[0][s]
	}

	for col := 1; col < m; col++ {
		for s := 0; s < k; s++ {
			ndp[s] = INF

			for t := 0; t < k; t++ {
				// 当前列状态 s 和上一列状态 t 必须兼容
				if compatible[t][s] {
					if dp[t]+cost[col][s] < ndp[s] {
						ndp[s] = dp[t] + cost[col][s]
					}
				}
			}
		}

		dp, ndp = ndp, dp
	}

	ans := INF
	for _, v := range dp {
		if v < ans {
			ans = v
		}
	}

	fmt.Fprintln(out, ans)
}
