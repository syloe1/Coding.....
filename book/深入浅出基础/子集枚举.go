/*
从 n个整数里面选 k 个数相加，求和是质数的方案总数。
n <20
循环：`mask` 从 `0` 到 `(1 << n) - 1`，遍历全部子集
`mask >> i & 1`：取出第 i 位，等于 1 → 选第 i 个数；0 → 不选
过滤条件：子集里面 **1 的个数 == k**（恰好选 k 个），然后求和，判断是不是质数
mask = 5 里面2个1
mask = 11 里面3个1
*/
package main

import "fmt"

func isprime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	total := 1 << n //1M
	for mask := 0; mask < total; mask++ {
		cnt := 0 // 选了几个
		sum := 0 //选数字的和
		//数组  5 7 9
		for i := 0; i < n; i++ {
			//mask 5 (101) 选到了5 9
			if (mask >> i & 1) == 1 {
				cnt++
				sum += a[i]
			}
		}
		if cnt == k && isprime(sum) {
			ans++
		}
	}
	fmt.Println(ans)
}
