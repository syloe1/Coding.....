/*
给你一个整数 n 。现有一个包含 n 个顶点的 无向 图，顶点按从 0 到 n - 1 编号。给你一个二维整数数组 edges 其中 edges[i] = [ai, bi] 表示顶点 ai 和 bi 之间存在一条 无向 边。

返回图中 完全连通分量 的数量。

如果在子图中任意两个顶点之间都存在路径，并且子图中没有任何一个顶点与子图外部的顶点共享边，则称其为 连通分量 。

如果连通分量中每对节点之间都存在一条边，则称其为 完全连通分量 。

*/
// 判断每个连通块的 边和点 满不满足 完全图定义 点n  * (n-1)/2 = 边
// 按照无向图做，判断n * (n-1) == 边
func countCompleteComponents(n int, edges [][]int) int {
	ans := 0
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vs := make([]bool, n)
	var v, e int
	var dfs func(int)
	dfs = func(x int) {
		vs[x] = true
		v++
		e += len(g[x])
		for _, y := range g[x] {
			if !vs[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vs {
		if !b {
			v, e = 0, 0
			dfs(i)
			if e == (v-1)*v {
				ans++
			}
		}
	}
	return ans
}