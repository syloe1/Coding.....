/*
题目：前缀和
输入一个长度为 n 的整数序列。
接下来再输入 m 个询问，每个询问输入一对 l,r。
对于每个询问，输出原序列中从第 l 个数到第 r 个数的和。

输入格式：
第一行包含两个整数 n 和 m。
第二行包含 n 个整数，表示整数数列。
接下来 m 行，每行包含两个整数 l 和 r，表示一个询问的区间范围。

输出格式：
共 m 行，每行输出一个询问的结果。

数据范围：
1 ≤ l ≤ r ≤ n
1 ≤ n,m ≤ 100000
−1000 ≤ 数列中元素的值 ≤ 1000
*/
// s[i] 代表前 i 个元素的总和
// sum(l, r) = s[r] - s[l-1]
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	first := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(first[0])
	m, _ := strconv.Atoi(first[1])
	sc.Scan()
	parts := strings.Split(sc.Text(), " ")
	a := make([]int, n+1) // a[1]~a[n]
	s := make([]int, n+1) // s[0]=0
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.Atoi(parts[i-1])
		s[i] = s[i-1] + a[i]
	}
	for i := 0; i < m; i++ {
		sc.Scan()
		lr := strings.Split(sc.Text(), " ")
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])
		fmt.Println(s[r] - s[l-1])
	}
}
