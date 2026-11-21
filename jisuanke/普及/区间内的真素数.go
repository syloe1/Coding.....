// https://www.jisuanke.com/problem/T1174
/*
蒜头君请你帮忙找出正整数 M 和 N 之间 (N 不小于 M) 的所有真素数。
真素数的定义: 如果一个正整数 P 为素数，且其反序也为素数，那么 P 就为真素数。
例如，11，13 均为真素数，因为 11 的反序还是为 11，13 的反序为 31 也为素数。

输入格式
输入两个数 M 和 N，空格间隔，1 ≤ M ≤ N ≤ 100000。

输出格式
按从小到大输出 M 和 N 之间（包括 M 和 N ）的真素数，逗号间隔。如果之间没有真素数，则输出 No。

格式说明
输出时每行末尾的多余空格，不影响答案正确性

样例输入
10 35

样例输出
11,13,17,31

*/
package main

import "fmt"

const MAX = 100000

func sieve(max int) []bool {
	prime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		prime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if prime[i] { //倍数肯定是合数
			for j := i * i; j <= max; j += i {
				prime[j] = false
			}
		}
	}
	return prime
}
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 { //x是合数
			return false
		}
	}
	return true
}
func reverseNum(x int) int {
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}
func main() {
	primeTable := sieve(MAX)
	var m, n int
	fmt.Scan(&m, &n)
	var ans []int
	for p := m; p <= n; p++ {
		if primeTable[p] {
			rp := reverseNum(p)
			var f bool
			if rp <= MAX {
				f = primeTable[rp]
			} else {
				f = isPrime(rp)
			}
			if f {
				ans = append(ans, p)
			}
		}
	}
	if len(ans) == 0 {
		fmt.Println("No")
	} else {
		for idx, v := range ans {
			if idx > 0 {
				fmt.Print(",")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
