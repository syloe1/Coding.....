/*
你正在维护一个项目，该项目有 n 个方法，编号从 0 到 n - 1。

给你两个整数 n 和 k，以及一个二维整数数组 invocations，其中 invocations[i] = [ai, bi] 表示方法 ai 调用了方法 bi。

已知如果方法 k 存在一个已知的 bug。那么方法 k 以及它直接或间接调用的任何方法都被视为 可疑方法 ，我们需要从项目中移除这些方法。

只有当一组方法没有被这组之外的任何方法调用时，这组方法才能被移除。

返回一个数组，包含移除所有 可疑方法 后剩下的所有方法。你可以以任意顺序返回答案。如果无法移除 所有 可疑方法，则 不 移除任何方法。



示例 1:

输入: n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]

输出: [0,1,2,3]
*/
/*
可疑方法是：k 本身，以及从 k 出发沿调用边能到达的所有方法。
移除条件不是“不能有人调用 k”，而是：
可疑集合里的任何一个方法，都不能被可疑集合之外的方法调用。
从 k 开始 DFS 图，标记所有可能有 bug 的方法（节点）。题目把这些方法叫做可疑方法。
遍历 invocations，如果存在从「非可疑方法」到「可疑方法」的边，则删除后无法编译，返回数组 [0,1,2,⋯,n−1]。
否则可以正常删除，把非可疑方法加入答案。
⚠注意：图中可能有环，为避免 DFS 无限递归下去，只需 DFS 没有访问过（没有被标记）的节点。转载请注明出处。
*/
func remainingMethods(n int, k int, invocations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, e := range invocations {
		g[e[0]] = append(g[e[0]], e[1])
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
	dfs(k)
	// vs[i] = true, x 是可疑方法
	for _, e := range invocations {
		if !vs[e[0]] && vs[e[1]] {
			for i := range n {
				ans = append(ans, i)
			}
			return
		}
	}
	for i, b := range vs {
		if !b {
			ans = append(ans, i)
		}
	}
	return
} 