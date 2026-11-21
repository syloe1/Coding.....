/*
约数个数定理：
如果一个大于 1 的正整数 n，唯一分解为 n = p1^a1 * p2^a2 * ... * pk^ak
p1,p2...pk 是互不相同质数，a1,a2...ak 是对应指数
那么 n 的正约数总个数 = (a1+1)*(a2+1)*...*(ak+1)
原理：每个质因子可以选 0 次到 a 次，一共 a+1 种选择，独立选择，乘法原理。
约数和定理：
n 的所有正约数之和 = (1+p1+p1^2+...+p1^a1) * (1+p2+p2^2+...+p2^a2) * ... * (1+pk+pk^2+...+pk^ak)
*/
package main

// Factor 唯一分解，返回 map[质数]指数
func Factor(n int64) map[int64]int64 {
	res := make(map[int64]int64)
	if n <= 1 {
		return res
	}
	// 处理2
	for n%2 == 0 {
		res[2]++
		n /= 2
	}
	// 处理奇数，j <= x/j 防止溢出
	var j int64 = 3
	for j <= n/j {
		if n%j == 0 {
			for n%j == 0 {
				res[j]++
				n /= j
			}
		}
		j += 2
	}
	if n > 1 {
		res[n]++
	}
	return res
}

// DivisorCount 约数个数定理
func DivisorCount(f map[int64]int64) int64 {
	ans := int64(1)
	for _, a := range f {
		ans *= (a + 1)
	}
	return ans
}

// DivisorSum 约数和定理
func DivisorSum(f map[int64]int64) int64 {
	ans := int64(1)
	for p, a := range f {
		sum := int64(1)
		term := int64(1)
		for i := int64(0); i < a; i++ {
			term *= p
			sum += term
		}
		ans *= sum
	}
	return ans
}
