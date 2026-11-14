/*
给你一个有 n 个节点的 有向无环图（DAG），请你找出从节点 0 到节点 n-1 的所有路径并输出（不要求按特定顺序）

	graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
*/
package main

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	n := len(graph)
	path := []int{}
	var dfs func(u int)
	dfs = func(u int) {
		path = append(path, u)
		if u == n-1 {
			ans = append(ans, append([]int(nil), path...))
		} else {
			for _, v := range graph[u] {
				dfs(v)
			}
		}
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}
