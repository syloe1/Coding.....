/*
给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。

请你返回 无法互相到达 的不同 点对数目 。
*/
/*
建图后，用 DFS 可以求出每个连通块的大小。

求连通块的大小的同时，用一个变量 total 维护前面求出的连通块的大小之和。

设当前连通块的大小为 size，那么这个连通块中的每个点，与前面遍历过的连通块的每个点，都是无法互相到达的，根据乘法原理，这有 size⋅total 个，加到答案中。
*/
func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}
	vs := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vs[x] = true
		size := 1
		for _, y := range g[x] {
			if !vs[y] {
				size += dfs(y)
			}
		}
		return size
	}
	total := 0
	for i, b := range vs {
		if !b {
			size := dfs(i)
			ans += int64(size) * int64(total)
			total += size
		}
	}
	return
}