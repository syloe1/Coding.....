/*
给你一个正整数 n ，表示总共有 n 个城市，城市从 1 到 n 编号。给你一个二维数组 roads ，其中 roads[i] = [ai, bi, distancei] 表示城市 ai 和 bi 之间有一条 双向 道路，道路距离为 distancei 。城市构成的图不一定是连通的。
两个城市之间一条路径的 分数 定义为这条路径中道路的 最小 距离。
返回城市 1 和城市 n 之间的所有路径的 最小 分数。
*/
/*
由于路径可以折返，节点 1 所在连通块的每条边都可以在路径上，所以答案为节点 1 所在连通块中的最小的 distance
i 。
*/
func minScore(n int, roads [][]int) int {
	type edge struct {
		to, dis int
	}
	g := make([][]edge, n+1)
	for _, e := range roads {
		x, y, dis := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, dis})
		g[y] = append(g[y], edge{x, dis})
	}
	vs := make([]bool, n+1)
	ans := math.MaxInt
	var dfs func(int)
	dfs = func(x int) {
		vs[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.dis)
			if !vs[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(1)
	return ans
}