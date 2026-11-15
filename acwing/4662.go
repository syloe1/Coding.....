/*
记 f(x) 为 x 的所有因数（约数）的平方的和。
例如：f(12)=1²+2²+3²+4²+6²+12²。
定义 g(n)=∑_{i=1}^n f(i)。
给定 n，求 g(n) 除以 10⁹+7 的余数。

输入格式
输入一行包含一个正整数 n。

输出格式
输出一个整数表示答案 g(n) 除以 10⁹+7 的余数。

数据范围
对于 20% 的评测用例，n≤10⁵。
对于 30% 的评测用例，n≤10⁷。
对于所有评测用例，1≤n≤10⁹。

输入样例：
100000
输出样例：
680584257

*/
/*
先把定义再复述一遍：
f(x)：x 的所有约数的平方相加。 比如 f(12)=1²+2²+3²+4²+6²+12²
g(n)=∑_{i=1}^n f(i)：把 f(1),f(2),f(3)…f(n) 全部加起来
第一步：展开 g(n)
g(n)=∑_{i=1}^n f(i)=∑_{i=1}^n (∑_{d|i} d²)
内层 ∑_{d|i}d² 的意思：枚举 i 的所有约数 d，累加 d²。
所以整个式子等价于：
遍历每一个 i（从 1 到 n），再找出 i 的每一个约数 d，把 d² 加进总和。
遍历每一个 i（从 1 到 n），再找出 i 的每一个约数 d，把 d² 加进总和。现在换个视角：不按 i 去枚举，反过来按 d 去枚举。 我们看：数字 d² 在总和里一共出现多少次？ 什么时候 d 会被算进去？只要 d 是某个 i 的约数，也就是 i 是 d 的倍数。
在 1~n 里面，d 的倍数是：
d, 2d, 3d,…, ⌊n/d⌋⋅d 总共有 ⌊n/d⌋ 个！ 每出现一次，就要加上一个 d²。 所以 d² 对总和的贡献就是：d² ⋅ ⌊n/d⌋
所有 d 从 1 到 n 全部加起来：
g(n)=∑_{d=1}^n d²⋅⌊n/d⌋
*/
/*
怎么快速算
直接枚举 d=1…n 是 O(n)，n=1e9 还是不行。
注意 ⌊n/d⌋ 只有 O(√n) 种不同取值，所以可以数论分块。
对于一段区间 [l,r]，如果：
⌊n/d⌋=q
那么这一段的贡献是：
q⋅∑_{d=l}^r d²
其中：
∑_{d=l}^r d² = S(r) − S(l−1)
平方和公式：
S(m)=∑_{i=1}^m i² = m(m+1)(2m+1)/6
因为模数是 1e9+7，除法要换成乘 6 的逆元。
*/
/*
交换求和
两层求和：
外层：i 从 1 枚举到 n
内层：对每一个 i，枚举它所有的约数 d，累加 \(d^2\)
固定 d，找所有合法 i。
什么时候 (i,d) 是合法的？条件为 d 整除 i，即 i 是 d 的倍数。i 的取值范围是 1 到 n，满足条件的 i 分别是 d、2d、3d 直到 kd，总共有 n 除以 d 向下取整个。
单个d贡献 dd * (n/d)
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int64 = 1_000_000_007

func powMod(a, b int64) int64 {
	res := int64(1)
	a %= mod
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

// 逆元 x 的模 p 乘法逆元等于 x 的 \(p-2\) 次方再对 p 取模。
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int64
	fmt.Fscan(in, &n)

	inv := powMod(6, mod-2)
	//s(m) = 1 ^ 2 + ... + m ^ 2
	sum := func(m int64) int64 {
		if m < 0 {
			return 0
		}
		mm := m % mod
		res := mm
		res = res * ((m + 1) % mod) % mod
		res = res * ((2*m + 1) % mod) % mod
		res = res * inv % mod
		return res
	}
	ans := int64(0)
	for l := int64(1); l <= n; {
		q := n / l
		r := n / q
		part := (sum(r) - sum(l-1) + mod) % mod
		ans = (ans + (q%mod)*part) % mod

		l = r + 1
	}
	fmt.Fprintln(out, ans)
}
