// https://www.jisuanke.com/problem/T1243
/*
不超过n位的正整数中，有多少个回文数 ? n <= 10
输出: 一个整数， 即回文数个数
*/
/*
 回文数计数公式
 k奇数  len = (k + 1) / 2, cnt = pow(9, 10 ^ (len-1))
 **k 位回文由前一半数字唯一确定**，第一位不能是 0
*/
package main

import "fmt"

func pow10(x int) int {
	res := 1
	for i := 0; i < x; i++ {
		res *= 10
	}
	return res
}
func countK(k int) int {
	t := (k - 1) / 2
	return 9 * pow10(t)
}
func main() {
	var n int
	fmt.Scan(&n)
	tot := 0
	for k := 1; k <= n; k++ {
		tot += countK(k)
	}
	fmt.Println(tot)
}
