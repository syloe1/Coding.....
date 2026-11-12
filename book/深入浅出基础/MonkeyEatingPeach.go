/*
 猴 子 吃 桃 。 一 只 小 猴 买 了 若 干 桃 子 。 第 一 天 它 吃 了 这 些 桃 子 的 一 半 之 后 ， 又 贪
嘴 多 吃 了 一 个 ， 第 二 天 它 也 吃 了 剩 余 桃 子 的 一 半 ， 又 贪 嘴 多 吃 了 一 个 ； 第 三 天 它 又 吃 了 剩 下 的 桃
子 的 一 半 ， 并 贪 嘴 多 吃 了 一 个 。 第 四 天 起 来 一 看 ， 发 现 桃 子 只 剩 下 一 个 了 。 请 问 ： 小 猴 买 了 几 个
桃 子
*/
// f[k] 第k天早上的桃子数
package main

import "fmt"

// days：还需要往前推几天
func dfs(days int) int {
	// 基线：推完3天，第四天剩余1
	if days == 0 {
		return 1
	}
	// 逆推公式：今天早上 = (明天早上 + 1) * 2
	next := dfs(days - 1)
	return (next + 1) * 2
}

func main() {
	res := dfs(3) // 需要往前推3天
	fmt.Println("初始桃子数量 =", res)
}
