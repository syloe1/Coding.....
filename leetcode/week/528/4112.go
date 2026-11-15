/*
给你一个长度为 n 的整数数组 parent，表示一棵节点编号从 0 到 n - 1 的有根树，其中 parent[i] 表示节点 i 的父节点。树的根节点为节点 0，因此 parent[0] = -1。
另给你两个长度均为 n 的整数数组 op 和 score，其中 op[i] 决定节点 i 允许执行的操作，score[i] 表示该节点的分数。
你需要按照以下规则处理每个 非叶节点 i：
如果 op[i] = 0，你可以选择节点 i 的 至多 一个子节点，并将以该子节点为根的 子树 中所有节点的分数设为 0。
如果 op[i] = 1，你必须选择节点 i 的 恰好 一个子节点，并将以该子节点为根的子树中所有节点的分数设为 0。
每个非叶节点都必须恰好处理一次。你可以按任意顺序处理这些节点。
返回执行这些操作后，所有节点分数之和的 最大可能值。
叶节点 是没有子节点的节点。
以某个节点为根的 子树 由该节点及其所有后代节点组成。
n < 1e5 p.length = op.length = score.length  = n
parent[0] = -1
score[i] 在[-1e4, 1e4]
*/
/*
为了总分最大化，我们肯定希望清零那些“得分为负数”的子树。
如果必须清零（op=1），即使子节点都是正数，也要“两害相权取其轻”，清零得分最小的那个。
树形 DP + 自底向上（拓扑排序 / BFS）。 用 BFS 模拟拓扑排序最稳妥
dp[u]：以节点 u 为根的子树，在内部执行合理的清零操作后，能提供的最大得分。
（前提假设：节点 u 本身没有被它的父节点清零，因为 dp[u] 是真实贡献给父节点的值）。
设节点 u 的子节点集合为 children[u]。
基础得分 base = score[u] + sum(dp[v])，其中 v 遍历所有子节点。

叶子节点 (len(children[u]) == 0)
dp[u] = score[u]
【避坑点】：绝对不能写成 max(0, score[u])。叶子节点没有权限清零自己，如果提前截断为0，父节点在做“必须清零一个”的决策时，就会丢失真实的负数代价，导致算出错误的偏大结果。

非叶子节点，且 op[u] == 1（必须清零恰好一个子节点）
dp[u] = base - min(dp[v])
（解释：选出得分最小的子节点清零，相当于扣除它原本能带来的贡献，这样损失最小。）

非叶子节点，且 op[u] == 0（至多清零一个子节点）
dp[u] = base - min(0, min(dp[v]))
（解释：如果最小的子节点是负数，就清零它，相当于加上了它的绝对值；如果是正数，就不清零。）
*/
package main

func maxScore(parent []int, op []int, score []int) int {
	n := len(parent)

	// 1. 建树
	children := make([][]int, n)
	childCount := make([]int, n)

	for i := 1; i < n; i++ {
		p := parent[i]
		// children[p] 记录的是 p的孩子节点
		children[p] = append(children[p], i)
		childCount[p]++ //记录每个节点有多少个孩子1
	}

	// 2. 初始化队列：所有叶子节点入队
	queue := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if childCount[i] == 0 {
			queue = append(queue, i)
		}
	}

	dp := make([]int, n)            //以 u 为根的子树，经过内部最优操作后，能够贡献的最大分数。
	sumChildrenDp := make([]int, n) //节点 u 所有已经处理完毕的孩子的 DP 值之和。

	head := 0

	// 3. 自底向上处理
	for head < len(queue) {
		u := queue[head]
		head++

		// 叶子节点
		if len(children[u]) == 0 {
			dp[u] = score[u]
		} else {
			base := score[u] + sumChildrenDp[u]

			// 找最小的子树贡献
			minChildDp := dp[children[u][0]]

			for _, v := range children[u][1:] {
				if dp[v] < minChildDp {
					minChildDp = dp[v]
				}
			}

			if op[u] == 1 {
				// 必须清零一个子树
				dp[u] = base - minChildDp
			} else {
				// 至多清零一个子树
				dp[u] = base
				if minChildDp < 0 {
					dp[u] -= minChildDp
				}
			}
		}

		// 4. 向父节点传递贡献
		p := parent[u]

		if p != -1 {
			sumChildrenDp[p] += dp[u]
			childCount[p]--

			if childCount[p] == 0 {
				queue = append(queue, p)
			}
		}
	}

	return dp[0]
}
