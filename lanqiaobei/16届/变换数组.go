/*
题目 3337: 蓝桥杯2025年第十六届省赛真题-变换数组 时间限制: 1s 内存限制: 128MB
题目描述
输入一个数组 a，包含有 n 个元素 a₁, a₂, ⋯, aₙ。对这个数组进行 m 次变换，每次变换会将数组 a 中的每个元素 aᵢ 转换为 aᵢ·bitCount(aᵢ)。其中 bitCount(x) 表示数字 x 的二进制表示中 1 出现的次数，例如 bitCount(3) = 2，因为 3 的二进制表示为 11，其中 1 出现了两次。
请输出变换之后的数组内容。

输入格式
输入的第一行包含一个正整数 n，表示数组 a 中的元素个数。
第二行包含 n 个整数 a₁, a₂, ⋯, aₙ，相邻整数之间使用一个空格分隔。
第三行包含一个整数 m，表示变换次数。

输出格式
输出一行，包含 n 个整数，相邻整数之间使用一个空格分隔，表示变换之后得到的数组 a。

样例输入
2
5 7
2
样例输出
20 63
对于所有评测用例，1 ≤ n ≤ 1000 ，0 ≤ m ≤ 5 ，0 ≤ ai ≤ 1000 。
*/

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var m int
	fmt.Fscan(in, &m)

	// 进行 m 次变换
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			count := bits.OnesCount(uint(a[i]))
			a[i] *= count
		}
	}

	// 输出结果
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, a[i])
	}
	fmt.Fprintln(out)
}
