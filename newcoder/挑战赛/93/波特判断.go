//https://ac.nowcoder.com/acm/contest/141878/A
/*
因为波特很喜欢加括号，所以本题给定一个长度为 n 的字符串 s = s1s2…sn，保证其仅包含小写字母和小括号字符 (、)。

你需要找出所有满足以下条件的下标 i (1 ≤ i < n)：

- 字符 si = (且 sn =)；
- 不存在整数 j (i < j < n)，使得 sj = ( 或 sj = )。

输入描述：
第一行输入一个整数 n (1 ≤ n ≤ 10^5)，表示字符串的长度。
第二行输入一个长度为 n 的字符串 s，保证其仅包含小写字母和小括号字符 (、)。

输出描述：
如果不存在合法的下标 i，输出 -1；
否则，按从小到大的顺序输出所有合法的下标 i。
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	if n < 2 || s[n-1] != ')' {
		fmt.Println(-1)
		return
	}
	pos := -1
	for i := n - 2; i >= 0; i-- {
		if s[i] == '(' || s[i] == ')' {
			pos = i
			break
		}
	}
	if pos == -1 {
		fmt.Println(-1)
	} else if s[pos] == '(' {
		fmt.Println(pos + 1)
	} else {
		fmt.Println(-1)
	}
}
