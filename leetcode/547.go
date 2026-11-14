/*
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。
n <= 200
*/
/*
dfs(u):
    visited[u] = true
    for v 所有邻居:
        if 连通 && 没访问:
            dfs(v)
DFS 递归：**递归契约 + base case + recursive case**
- base case：归纳起点，不再拆分
- recursive case：依赖子问题的正确性（归纳假设）
- 契约：这个函数承诺完成什么，保证归纳链条不断
*/
// 遍历连通块 + 打标记
//连通块染色 DFS
package main

func findCircleNum(isConnected [][]int) (ans int) {
	vs := make([]bool, len(isConnected))
	var dfs func(u int)
	dfs = func(u int) {
		vs[u] = true
		for v := 0; v < len(isConnected); v++ {
			if isConnected[u][v] == 1 && !vs[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < len(isConnected); i++ {
		if !vs[i] {
			ans++
			dfs(i)
		}
	}
	return
}
