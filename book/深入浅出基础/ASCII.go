package main

import "fmt"

func main() {
	count := 0
	for c := byte('A'); c <= 'Z'; c++ {
		fmt.Printf("%c ", c)
		count++
		// 每5个换行
		if count%5 == 0 {
			fmt.Println()
		}
	}
	// 1) M 是第几个
	var c byte = 'K'
	idx := c - 'A' + 1
	fmt.Println(idx)

	// 2) 第18个字母是什么
	n := 18
	ch := 'A' + (n - 1)
	fmt.Println(string(ch))
}
