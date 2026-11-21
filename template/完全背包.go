/*
区别 01 背包：**每种物品可以拿无数次**（只要背包装得下），不是选 / 不选二选一，可以选 0、1、2、3…… 多个。
先二维 DP，再一维优化，和 01 背包对照着学，重点对比转移和遍历顺序。
背包容量 V，一共 n 个物品
物品 i：重量 \(w[i]\)，价值 \(v[i]\)
物品不限数量，可以反复拿。
求：不超重下最大总价值。
dp[i][j] 只考虑前 i 种物品，背包容量 j 的最大价值
   不选 第i种物品，  dp[i][j] = dp[i-1][j]
   选 第i种物品  dp[i][j] = dp[i][j-w[i]] +  v[i]
*/
// TwoDimComplete 完全背包二维DP
// n:物品种类数 V:背包容量
// w,v下标从1开始
package main

func TwoDimComplete(n, V int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, V+1)
	}
	// 枚举物品
	for i := 1; i <= n; i++ {
		for j := 1; j <= V; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-w[i]]+v[i])
			}
		}
	}
	return dp[n][V]
}

/*
优化 二维依赖：`dp[i][j]` 依赖 `dp[i-1][j]`（上一行） 和 `dp[i][j-w[i]]`（**当前行左边，已经算过的值**）
dp[j] 容量j的最大价值
dp[j] = max(dp[j], dp[j-w[i]] + v[i])
j正序 `j-w[i] < j`，正序的时候，`dp[j-w[i]]` 在本轮 i 已经更新过，就是二维里 `dp[i][j-w[i]]` 的含义，代表还可以继续拿当前物品。
*/

/*
- 01 背包一维：j **逆序**，用 i-1 旧值，物品只能一次
- 完全背包一维：j **正序**，用 i 本轮新值，物品无限
*/
func OneDimComplete(n, V int, w, v []int) int {
	dp := make([]int, V+1)
	for i := 1; i <= n; i++ {
		// 正序
		for j := w[i]; j <= V; j++ {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}
	return dp[V]
}
