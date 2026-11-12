/*
楼梯有 N 阶，上楼可以一步上一阶，也可以一步上二阶。计算共有多少种不同的走法。

**输入**：一个正整数 N（\(1\le N \le 5000\)）
**输出**：一个正整数，表示走法总数
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n+1)
	s[0], s[1] = 1, 1
	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	fmt.Println(s[n])
}
