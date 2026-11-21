/*
给你一个整数 n，表示一个 有向加权 图中的节点数量，节点编号从 0 到 n - 1。该图由二维数组 edges 表示，其中 edges[i] = [ui, vi, wi] 表示一条从节点 ui 指向节点 vi、权重为 wi 的有向边。
另给定一个长度为 n 的字符串 labels，其中 labels[i] 是分配给节点 i 的字符，以及一个整数 k。
返回一条从节点 0 到节点 n - 1 的路径的 最小总边权 ，并要求该路径上所有节点标签按顺序 拼接 后，最多包含 k 个 连续相同 字符。如果不存在有效路径，返回 -1。
注意：节点和边可能会被重复访问。
n <= 5e4
wi < 1e4
k < 50
*/
/*
题目要求：
- 从节点 0 到节点 n-1。
- 路径总边权最小。
- 路径上连续相同的字符不能超过 k 个。
- 节点和边允许重复访问。
如何设计状态？
我们需要思考：为了判断下一步能不能走，除了当前节点，还需要知道什么？

显然需要知道：
1. 当前所在节点 u。
2. 当前连续相同字符出现了多少次 cnt。
注意：当前最后一个字符不需要额外记录，因为它就是 labels[u]。
state = (node, cnt)
过一条边（u， v)
字符相同 newcnt := cnt + 1
		if newcnt > k 跳过
		else 进入 (v, newcnt)
字符不同  newcnt := 1 进入 (v, 1)
	newdist = dist + w
答案是 min{(n-1,1)...(n-1,k)}
*/
import "container/heap"

type Edge struct {
	to     int
	weight int
}

type State struct {
	node int
	cnt  int
	dist int
}

type MinHeap []State

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(State))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	// 1. 建立邻接表
	graph := make([][]Edge, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{
			to:     v,
			weight: w,
		})
	}

	// 2. dist[u][cnt] 表示到达节点 u，
	//    且当前连续相同字符数量为 cnt 时的最短距离
	const INF = int(^uint(0) >> 1)

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, k+1)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	// 3. 初始化起点
	// 起点本身算一个字符，所以 cnt = 1
	dist[0][1] = 0

	h := &MinHeap{}
	heap.Init(h)

	heap.Push(h, State{
		node: 0,
		cnt:  1,
		dist: 0,
	})

	// 4. Dijkstra
	for h.Len() > 0 {
		cur := heap.Pop(h).(State)

		u := cur.node
		cnt := cur.cnt
		d := cur.dist

		// 跳过过期状态
		if d > dist[u][cnt] {
			continue
		}

		// 到达终点，当前状态已经是全局最优
		if u == n-1 {
			return d
		}

		// 5. 遍历当前节点的所有出边
		for _, e := range graph[u] {
			v := e.to
			w := e.weight

			newCnt := 1

			if labels[u] == labels[v] {
				newCnt = cnt + 1
			}

			// 连续相同字符超过 k，不合法
			if newCnt > k {
				continue
			}

			newDist := d + w

			// 松弛操作
			if newDist < dist[v][newCnt] {
				dist[v][newCnt] = newDist

				heap.Push(h, State{
					node: v,
					cnt:  newCnt,
					dist: newDist,
				})
			}
		}
	}

	return -1
}