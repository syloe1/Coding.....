//https://www.jisuanke.com/problem/T2029

/*
有两个仅包含小写英文字母的字符串A和B。现在要从字符串 A 中取出k个互不重叠的非空子串，然后把这 k 个子串按照其在字符串 A 中出现的顺序依次连接起来得到一个新的字符串，请问有多少种方案可以使得这个新串与字符串 B 相等？注意：子串取出的位置不同也认为是不同的方案。

输入格式
第一行是三个正整数 n,m,k，分别表示字符串 A 的长度，字符串 B 的长度，以及问题描述中所提到的 k，每两个整数之间用一个空格隔开。
第二行包含一个长度为 n 的字符串，表示字符串A。
第三行包含一个长度为 m 的字符串，表示字符串B。

输出格式
输出共一行，包含一个整数，表示所求方案数。由于答案可能很大，所以这里要求输出答案对 1,000,000,007 取模的结果。

数据范围
n <= 1000, m <= 200, k <= m
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
