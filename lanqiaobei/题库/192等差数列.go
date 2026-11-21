//https://www.lanqiao.cn/problems/192/learning/?page=1&first_category_id=1&name=%E7%AD%89%E5%B7%AE%E6%95%B0%E5%88%97
/*
数学老师给小明出了一道等差数列求和的题目。但是粗心的小明忘记了一部分的数列，只记得其中 N 个整数。
现在给出这 N 个整数，小明想知道包含这 N 个整数的最短的等差数列有几项。

【输入格式】
输入的第一行包含一个整数 N。
第二行包含 N 个整数 A_1,A_2,…,A_N。(注意 A_1 ∼ A_N 并不一定是按等差数列中的顺序给出)
其中，2 ≤ N ≤ 10^5,0 ≤ A_i ≤ 10^9。

【输出格式】
输出一个整数表示答案。
*/
// gcd(0, x) = x
/*
等差数列里，**任意两个数的差一定是公差 d 的整数倍**。
反过来：我们手里一堆已经在等差数列里的数，它们两两之间的差，都必须是 d 的倍数。
所以 d 必须是所有这些差值的**公约数**；想要数列最短，就要让 d 尽可能大 → 取**最大公约数 gcd**。
*/
package main

import (
	"fmt"
	"sort"
)

// 欧几里得算法求最大公约数
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	// 排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	var d int64 = 0
	for i := 1; i < n; i++ {
		delta := a[i] - a[i-1]
		d = gcd(d, delta)
	}

	if d == 0 {
		// 所有数字相同
		fmt.Println(1)
	} else {
		// 末尾比首部多加 (n-1)d, 所以+1是项目
		ans := (a[n-1]-a[0])/d + 1
		fmt.Println(ans)
	}
}
