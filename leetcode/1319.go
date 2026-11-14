/*

用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。
*/
/*
n 个节点要连通，至少需要 n-1 条边（树的性质）。

如果 connections 的边数 len(connections) < n-1，总边数不够，无论怎么拔插都不可能连通所有计算机，直接返回 -1。
为什么边数够时答案是“连通分量数 - 1”

把当前网络划分成若干个连通分量，设数量为 components。

每个连通分量内部已经连通，要把 components 个分量连成一个整体，至少需要 components - 1 条边。

每条操作可以拔掉一条多余的边，再用来连接两个不同分量，一次操作减少一个分量，所以最少操作次数就是 components - 1。

为什么多余的边一定够：每个连通分量内部至少是一棵树，否则有环，环上多出来的边可以拔掉。把所有分量内部多余的边拔掉后，剩余可用边数至少是 components - 1，正好够连接所有分量。

结论：只要 len(connections) >= n-1，答案就是 components - 1；否则返回 -1。
*/
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	g := make([][]int, n)
	for _, e := range connections {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	vs := make([]bool, n)
	c := 0
	var dfs func(int)
	dfs = func(u int) {
		vs[u] = true
		for _, v := range g[u] {
			if !vs[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !vs[i] {
			c++
			dfs(i)
		}
	}
	return c - 1
}