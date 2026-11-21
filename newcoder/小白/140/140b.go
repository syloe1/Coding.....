/*
链接：https://ac.nowcoder.com/acm/contest/143195/B
来源：牛客网

有 n 件快递需要邮寄，编号为 1 到 n。它们的重量用数组 a₁,a₂,…,aₙ 表示，体积用数组 b₁,b₂,…,bₙ 表示，其中第 i 件快递的重量为 aᵢ，体积为 bᵢ。一件重量为 x、体积为 y 的快递，其邮费为 max(x,y)。
你可以将这些快递任意划分成若干组，每组内的快递合并成一件后邮寄。合并后的重量为组内所有快递的重量之和，体积为组内所有快递的体积之和。

请你计算将全部快递寄出所需的最小总邮费。

输入描述:
每个测试文件均包含多组测试数据。第一行输入一个整数 T(1≤T≤10⁴) 代表数据组数，每组测试数据描述如下：
第一行输入一个整数 n(1≤n≤2×10⁵)，表示快递数量。
第二行输入 n 个整数 a₁,a₂,…,aₙ(1≤aᵢ≤1000)，表示各件快递的重量。
第三行输入 n 个整数 b₁,b₂,…,bₙ(1≤bᵢ≤1000)，表示各件快递的体积。

除此之外，保证单个测试文件中所有测试数据的 n 之和不超过 2×10⁵。

输出描述:
对于每组测试数据，新起一行。
输出一个整数，表示将全部快递寄出的最小总邮费。

*/
/*
# max(a+c, b+d) ≤ max(a,b)+max(c,d)

记：
x = max(a,b), y = max(c,d)
由最大值定义：
a ≤ x, b ≤ x,     c ≤ y, d ≤ y
于是：
a+c ≤ x+y,   b+d ≤ x+y
a+c 和 b+d 这两个数都小于等于 x+y，那它们里面最大的那一个当然也 ≤ x+y：
max(a+c, b+d) ≤ x+y = max(a,b)+max(c,d)

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}
		var sa, sb int64
		for i := 0; i < n; i++ {
			sa += int64(a[i])
			sb += int64(b[i])
		}
		var ans int64
		if sa > sb {
			ans = sa
		} else {
			ans = sb
		}
		fmt.Fprintln(out, ans)
	}
}
