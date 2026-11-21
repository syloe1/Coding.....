//https://www.lanqiao.cn/problems/1020/learning/?page=1&first_category_id=1&name=%E9%98%B6%E4%B9%98%E7%BA%A6%E6%95%B0
/*
100 阶乘有多少个约数
*/
// 唯一分解定理
// 任何一个大于 1 的正整数 n，都可以唯一分解成若干质数的乘积：
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := make([]int, n+10)

	for i := 1; i <= n; i++ {
		x := i
		for j := 2; j <= x/j; j++ { // ✅ 改成 j <= x/j，不再 j*j 溢出
			if x%j == 0 {
				for x%j == 0 {
					x /= j
					cnt[j]++
				}
			}
		}
		if x > 1 {
			cnt[x]++
		}
	}

	ans := int64(1)
	for i := 1; i <= n; i++ {
		if cnt[i] != 0 {
			/*
				约数个数公式：把每一个指数 + 1，全部乘起来
				为什么指数要加 1：
				一个数分解后有质因子 p 的 a 次方。
				构造约数的时候，这个质因子 p 可以选 0 次、1 次、2 次…… 一直到 a 次。
			*/
			ans *= int64(cnt[i] + 1)
		}
	}
	fmt.Println(ans)
}
