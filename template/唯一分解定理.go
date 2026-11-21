// 任何一个大于 1 的正整数 n，都可以唯一分解成若干质数的乘积：
// Factor 唯一分解，返回 map[prime]exponent
// 对数字 `n` 做质因数分解
func Factor(n int64) map[int64]int {
	res := make(map[int64]int)
	if n <= 1 {
		return res
	}

	// 先处理2
	for n%2 == 0 {
		res[2]++
		n /= 2
	}

	// 处理奇数，i从3到sqrt(n)
	var i int64 = 3
	max := int64(math.Sqrt(float64(n)))
	for i <= max && n > 1 {
		for n%i == 0 {
			res[i]++
			n /= i
			max = int64(math.Sqrt(float64(n)))
		}
		i += 2
	}

	// 剩余大于1，是质数
	if n > 1 {
		res[n]++
	}
	return res
}
