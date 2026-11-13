/*
链接：https://ac.nowcoder.com/acm/contest/143195/A
来源：牛客网

小花有 n 颗糖果，你需要将它们装入编号为 1,2,3,… 的礼盒。奇数编号的礼盒最多装 k 颗糖果，偶数编号的礼盒最多装 k+1 颗糖果。
礼盒必须从 1 号开始按编号顺序使用，不能跳过礼盒。除最后一个使用的礼盒外，其余使用的礼盒都必须装满。

请你计算装下全部糖果最少需要使用多少个礼盒。

输入描述:
第一行输入两个整数 n,k (1≤n,k≤2×10^5)，表示糖果数量、礼盒容量参数。

输出描述:
输出一个整数，表示礼盒的最小使用数量。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)
	g := 2*k + 1
	q := n / g
	rem := n % g
	ans := 2 * q
	if rem > 0 {
		if rem <= k {
			ans += 1
		} else {
			ans += 2
		}
	}
	fmt.Println(ans)
}
