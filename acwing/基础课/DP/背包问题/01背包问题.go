/*
有 N 件物品和一个容量是 V 的背包。每件物品只能使用一次。
第 i 件物品的体积是 vi，价值是 wi。
求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。
接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。

数据范围
0 < N,V ≤ 1000
0 < vi,wi ≤ 1000
*/
package main

import "fmt"

func OneTim(n, V int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, V+1)
	}
	// 枚举物品
	for i := 1; i <= n; i++ {
		for j := 1; j <= V; j++ {
			if j < w[i] {
				// 放不下第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	return dp[n][V]
}
func main() {
	var n, V int
	fmt.Scan(&n, &V)
	a := make([]int, n+1)
	v := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i], &v[i])
	}
	ans := OneTim(n, V, a, v)
	fmt.Println(ans)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
