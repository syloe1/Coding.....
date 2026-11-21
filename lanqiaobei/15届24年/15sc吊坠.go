/*
小蓝想制作一个吊坠，他手上有 n 个长度为 m 的首尾相连的环形字符串{s1, s2, · · · , sn} ，他想用 n − 1 条边将这 n 个字符串连接起来做成吊坠，要求所有的字符串连完后形成一个整体。连接两个字符串 si, sj 的边的边权为这两个字符串的最长公共子串的长度（可以按环形旋转改变起始位置，但不能翻转），小蓝希望连完后的这 n − 1 条边的边权和最大，这样的吊坠他觉得最好看，请计算最大的边权和是多少。

输入格式
输入的第一行包含两个正整数 n, m ，用一个空格分隔。

接下来 n 行，每行包含一个长度为 m 的字符串，分别表示 s1, s2, · · · , sn 。

输出格式
输出一行包含一个整数表示答案。

题目：n 个点，要选 n-1 条边连通所有点，边权是两个环串的最长公共子串，求**边权总和最大** → 标准最大生成树（Kruskal / Prim）

>
> 为什么是最大生成树：连通图、n 点 n-1 边、总和最大，就是最大生成树定义。最小生成树是代价最小，这个反过来，边权越大越优先选。

## 子问题：两个环形字符串，求最长公共子串（可以旋转，不能翻转）

环形串技巧：`s + s`，所有旋转结果都是 `s+s` 的长度为 m 的子串。
比如 `aabb` → `aabbaabb`，所有旋转都在里面。
然后求两个串（s+s、t+t）之间的**最长公共子串，长度不能超过 m**（因为原串只有 m 长）。
m≤50，很小，暴力 DP 求最长公共子串完全没问题。
最长公共子串 DP：`dp[i][j]` = 以 s [i-1], t [j-1] 结尾的公共子串长度
if s[i-1]==t[j-1]: dp[i][j]=dp[i-1][j-1]+1
else dp[i][j]=0
Kruskal 策略：**贪心，优先拿权值最大的边**
步骤：

1. 把所有边，从大到小排序
2. 一条一条拿边：
   - 如果这条边连接的两个点**本来就不在同一个连通块** → 收下这条边，累加权值，把两个块合并
   - 如果两个点**已经在一块了** → 扔掉这条边（再拿就会形成环！我们不能有环，生成树不能有环）
3. 收到 `n-1` 条边就停，结束
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func lccs(s, t string, m int) int {
	a := s + s
	b := t + t
	la, lb := len(a), len(b)
	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}
	maxLen := 0
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					if maxLen == m {
						return m
					}
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	if maxLen > m {
		maxLen = m
	}
	return maxLen
}

type Edge struct {
	u, v, w int
}
type DSU struct {
	fa []int
}

func NewDSU(n int) *DSU {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return &DSU{fa: fa}
}
func (d *DSU) Find(x int) int {
	if d.fa[x] != x {
		d.fa[x] = d.Find(d.fa[x])
	}
	return d.fa[x]
}
func (d *DSU) Union(x, y int) bool {
	x = d.Find(x)
	y = d.Find(y)
	if x == y {
		return false
	}
	d.fa[y] = x
	return true
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &strs[i])
	}

	// 建边
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			w := lccs(strs[i], strs[j], m)
			edges = append(edges, Edge{u: i, v: j, w: w})
		}
	}

	// Kruskal 最大生成树：边从大到小排序
	sort.Slice(edges, func(p, q int) bool {
		return edges[p].w > edges[q].w
	})

	dsu := NewDSU(n)
	ans := 0
	cnt := 0

	for _, e := range edges {
		if dsu.Union(e.u, e.v) {
			ans += e.w
			cnt++
			if cnt == n-1 {
				break
			}
		}
	}

	fmt.Fprintln(out, ans)
}
