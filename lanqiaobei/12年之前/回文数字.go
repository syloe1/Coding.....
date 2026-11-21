/*
题目描述
观察数字：12321，123321  都有一个共同的特征，无论从左到右读还是从右向左读，都是相同的。这样的数字叫做：回文数字。

本题要求你找到一些5位或6位的十进制数字。满足如下要求：
该数字的各个数位之和等于输入的整数。
输入格式
一个正整数  n  (10< n< 100),  表示要求满足的数位和。
输出格式
若干行，每行包含一个满足要求的5位或6位整数。
数字按从小到大的顺序排列。
如果没有满足条件的，输出：-1
*/
package main

import "fmt"

func is(x int) bool {
	o := x
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return o == r
}
func digitsum(x int) int {
	s := 0
	for x > 0 {
		s += x % 10
		x /= 10
	}
	return s
}
func main() {
	var n int
	fmt.Scan(&n)
	f := false
	for i := 0; i <= 999999; i++ {
		if is(i) && digitsum(i) == n {
			fmt.Println(i)
			f = true
		}
	}
	if !f {
		fmt.Println(-1)
	}
}
