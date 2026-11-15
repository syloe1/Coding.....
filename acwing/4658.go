/*
给定正整数 n，请问有多少个质数是 n 的约数。

输入格式
输入的第一行包含一个整数 n。

输出格式
输出一个整数，表示 n 的质数约数个数。

数据范围
对于 30% 的评测用例，1≤n≤10000。
对于 60% 的评测用例，1≤n≤10^9。
对于所有评测用例，1≤n≤10^16。

输入样例：
396
输出样例：
3
样例解释
396 有 2,3,11 三个质数约数。

*/
/*
质因数分解  每找到一个质因子， 都要除尽 while (n % 2 == 0) n /= 2;
循环条件要用 i <= n / i，不要固定成原数的平方根
因为 n 在不停变小。
例如原来 n=1e16，除着除着可能很快变小。
最后如果 n > 1，它本身也是一个质因子
因为所有小于等于 sqrt(n) 的因子都已经试过了，剩下的一定是质数。
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
	var n int64
	fmt.Fscan(in, &n)
	if n == 1 {
		fmt.Fprintln(out, 0)
		return
	}
	ans := 0
	if n%2 == 0 {
		ans++
		for n%2 == 0 {
			n /= 2
		}
	}
	// 枚举奇数
	for i := int64(3); i <= n/i; i += 2 {
		if n%i == 0 {
			ans++
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		ans++
	}
	fmt.Fprintln(out, ans)
}
