/*
Ignatius likes to write words in reverse way. Given a single line of text which is written by Ignatius, you should reverse all the words and then output them.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scan(&n)
	scanner.Scan() // 吃掉换行，和原题getchar()作用一样

	for ; n > 0; n-- {
		line := scanner.Text()
		stack := make([]rune, 0)
		for _, ch := range line {
			if ch == ' ' {
				// 碰到空格，输出栈内单词
				for len(stack) > 0 {
					fmt.Printf("%c", stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				fmt.Print(" ")
			} else {
				stack = append(stack, ch)
			}
		}
		// 处理行末尾最后一个单词（后面没有空格！）
		for len(stack) > 0 {
			fmt.Printf("%c", stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		fmt.Println()
	}
}
