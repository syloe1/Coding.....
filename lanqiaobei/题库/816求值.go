//https://www.lanqiao.cn/problems/816/learning/?page=1&first_category_id=1&name=%E6%B1%82%E5%80%BC
/*
学习了约数后，小明对于约数很好奇，他发现，给定一个正整数 t，总是可以找到含有 t 个约数的整数。小明对于含有 t 个约数的最小数非常感兴趣，并把它定义为 S_t。
例如 S₁ = 1, S₂ = 2, S₃ = 4, S₄ = 6,…。


现在小明想知道，当 t = 100 时，S_t 是多少？即 S₁₀₀ 是多少？
*/

/*
质数从小到大依次取.
指数按**从大到小排好**，大指数优先分给更小的质数，这样整体乘积才最小。
约数个数公式：
如果 n = p₁^a₁ × p₂^a₂ × p₃^a₃ × p₄^a₄
那么 n 的约数总数 = (a₁+1)(a₂+1)(a₃+1)(a₄+1)
*/
/*
：把 t 拆成若干大于 1 的整数相乘，拆分的因子降序排列
100 = 5 * 5 *2 *2 = (a1 +1) * (a2 + 1) * (a3 + 1) * (a4 +1)
	指数 4 4 1 1
	用质数去构造 n =  (2 ^ 4 * 3 ^ 4 * 5 ^ 1 * 7 ^ 1)
*/
package main

import (
	"fmt"
)

// 预选质数，足够做题用：2,3,5,7,11,13,17,19,23,29,31
var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
var ans int64 // 存最终最小答案

// dfs 回溯
// pos: 当前用到第几个质数
// last: 上一个因子，保证因子降序（避免重复拆分）
// now: 当前已经算出的数
// cnt: 当前约数个数乘积
// target: 目标t
func dfs(pos int, last int64, now int64, cnt int64, target int64) {
	if cnt == target {
		if now < ans {
			ans = now
		}
		return
	}
	if pos >= len(primes) || cnt > target {
		return
	}

	// k 就是 (a+1)，k >=2，不能超过last，保证降序
	for k := int64(2); k <= last; k++ {
		if target%k != 0 {
			continue
		}
		// 指数 a = k - 1
		//k=4, a = 4
		next := now
		ok := true
		// 连续 乘 k-1次
		for i := int64(0); i < k-1; i++ {
			next *= primes[pos]
			if next > ans { // 剪枝：已经比当前答案大，不用继续乘了
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		// 下一个质数，下一个因子不能超过k
		dfs(pos+1, k, next, cnt*k, target)
	}
}

// pos: 当前用到第几个质数
// last: 上一个因子，保证因子降序（避免重复拆分）
// now: 当前已经算出的数
// cnt: 当前约数个数乘积
// target: 目标t
func findMinNumber(t int64) int64 {
	ans = 1e18 // 初始化成一个很大的数
	dfs(0, t, 1, 1, t)
	return ans
}

func main() {
	// 修改这里的t，就可以算 S_t
	var t int64
	fmt.Print("请输入t：")
	fmt.Scan(&t)
	res := findMinNumber(t)
	fmt.Printf("S_%d = %d\n", t, res)
}
