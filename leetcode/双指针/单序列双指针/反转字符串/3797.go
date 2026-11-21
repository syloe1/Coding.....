/*
给你一个字符串 s 和一个整数 k。

反转 s 的前 k 个字符，并返回结果字符串。
*/
func reversePrefix(s string, k int) string {
	/*
		库函数
		slices.Reverse(b[:k])
	*/
	b := []byte(s)
	if k > len(b) {
		k = len(b)
	}
	l, r := 0, k-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}