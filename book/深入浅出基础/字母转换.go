// 输 人 一 个 小 写 字 母 ， 输 出 其 对 应 的 大 写 字 母
package main

import "fmt"

func main() {
	var c byte
	fmt.Scanf("%c", &c)
	upper := c - 'a' + 'A'
	fmt.Printf("%c\n", upper)
}
