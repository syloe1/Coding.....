/*
棋盘上 A 点 (0,0) 有一个过河卒，需要走到右下角目标 B 点 (n,m)。卒行走规则：只能向下、或者向右走一格。
棋盘上 C 点有一匹对方的马，马所在位置，以及马跳一步能到达的 8 个点，都是控制点，卒不能踏入这些控制点。
求卒从起点 (0,0) 走到终点 (n,m) 的合法路径总数。

>
> 说明：马固定不动，不是走一步换位置。

## 输入格式

一行四个整数：n,m,x,y

- n,m：终点 B 的坐标 (n,m)
- x,y：马所在 C 点坐标 (x,y)
范围：1 ≤ n,m ≤ 20

## 输出格式

一个整数，合法路径总条数。
*/
// DP   dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[i][j] 走到坐标(i, j)的路径总数
package main

import "fmt"

func main() {
	var n, m, cx, cy int
	fmt.Scan(&n, &m, &cx, &cy)

	const flag = -1
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
	}
	dx := []int{2, 2, 1, 1, -1, -1, -2, -2}
	dy := []int{1, -1, 2, -2, 2, -2, 1, -1}

	dp[cx][cy] = flag
	for i := 0; i < 8; i++ {
		nx, ny := cx+dx[i], cy+dy[i]
		if nx >= 0 && ny >= 0 && nx <= n && ny <= n {
			dp[nx][ny] = flag
		}
	}
	if dp[0][0] == flag {
		fmt.Println(0)
		return
	}
	dp[0][0] = 1
	// 第一行
	for j := 1; j <= m; j++ {
		if dp[0][j] == flag {
			continue
		}
		if dp[0][j-1] == flag {
			dp[0][j] = 0
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// 初始化第一列
	for i := 1; i <= n; i++ {
		if dp[i][0] == flag {
			continue
		}
		if dp[i-1][0] == flag {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if dp[i][j] == flag {
				continue
			}
			var sum int64 = 0
			if dp[i-1][j] != flag {
				sum += dp[i-1][j]
			}
			if dp[i][j-1] != flag {
				sum += dp[i][j-1]
			}
			dp[i][j] = sum
		}
	}
	fmt.Println(dp[n][m])
}
