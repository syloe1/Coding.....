/*
给你一个正整数 n ，它表示一个 有向无环图 中节点的数目，节点编号为 0 到 n - 1 （包括两者）。

给你一个二维整数数组 edges ，其中 edges[i] = [fromi, toi] 表示图中一条从 fromi 到 toi 的单向边。

请你返回一个数组 answer，其中 answer[i]是第 i 个节点的所有 祖先 ，这些祖先节点 升序 排序。

如果 u 通过一系列边，能够到达 v ，那么我们称节点 u 是节点 v 的 祖先 节点。
*/

func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[y] = append(g[y], x)
	}
	vs := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vs[x] = true
		for _, y := range g[x] {
			if !vs[y] {
				dfs(y)
			}
		}
	}
	ans := make([][]int, n)
	for i := range ans {
		clear(vs)
		dfs(i)
		vs[i] = false
		for j, b := range vs {
			if b {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}