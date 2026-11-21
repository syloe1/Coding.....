// https://www.lanqiao.cn/problems/1037/learning/?page=1&first_category_id=1&name=%E5%AD%90%E4%B8%B2%E5%88%86%E5%80%BC%E5%92%8C
/*
【问题描述】
对于一个字符串 S，我们定义 S 的分值 f(S) 为 S 中出现的不同的字符个数。例如 f(aba)=2，f(abc)=3，f(aaa)=1。
现在给定一个字符串 S[0...n-1]（长度为 n），请你计算对于所有 S 的非空子串 S[i...j](0 ≤ i ≤ j < n)，f(S[i...j]) 的和是多少。

【输入格式】
输入一行包含一个由小写字母组成的字符串 S。
其中，1 ≤ n ≤ 10^5。

【输出格式】
输出一个整数表示答案。
*/
/*
核心思路：贡献法（逆向求和）
原式：ans = sum(每个子串的不同字符数)
等价转换：ans = sum(每个字符c，包含c的子串数量)
原理：交换求和顺序，避免枚举O(n²)个子串

计算单个字符贡献：
对于字符c在pos位置，上一次出现位置last[c]
新增合法子串数量 = (pos - last[c]) * (n - pos)
含义：
起点可以选 (last[c]+1 ~ pos) 一共 pos-last[c] 种
终点可以选 (pos ~ n-1) 一共 n-pos 种
保证子串包含当前这个c，且不会重复统计同一个子串里多次出现的c
*/
// 关键思路
/*
### 单个字符 c 怎么算它的贡献？

举例子：字符串 `abaca`，看字符 `a`
我们记录每个 `a` 上一次出现的位置 `last`
当前位置是 `pos`：

- 左边可选起点数量：\(pos - last\)
- 右边可选终点数量：\(n - pos\)
这个 a 带来的新增子串数量 = \((pos-last) \times (n-pos)\)

>
> 原理：保证这个子串**包含当前这个 a，并且不包含上一个 a**，避免重复统计！
*/
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	// last[c] 记录字符c上次出现的下标， 初始-1
	last := make([]int, 26)
	for i := range last {
		last[i] = -1
	}
	var ans int64 = 0
	for pos := 0; pos < n; pos++ {
		c := s[pos] - 'a'
		l := pos - last[c]
		r := n - pos
		ans += int64(l) * int64(r)
		last[c] = pos
	}
	fmt.Println(ans)
}
