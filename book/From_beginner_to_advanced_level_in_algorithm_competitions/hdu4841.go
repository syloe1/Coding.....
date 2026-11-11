/*
圆桌上围坐着2n个人。其中n个人是好人，另外n个人是坏人。如果从第一个人开始数数，数到第m个人，则立即处死该人；然后从被处死的人之后开始数数，再将数到的第m个人处死……依此方法不断处死围坐在圆桌上的人。试问预先应如何安排这些好人与坏人的座位，能使得在处死n个人之后，圆桌上围坐的剩余的n个人全是好人。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		n, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		table := make([]int, 0, 2*n)
		for i := 0; i < 2*n; i++ {
			table = append(table, i)
		}
		cur := 0
		//tick n people
		for i := 0; i < n; i++ {
			cur = (cur + m - 1) % len(table)
			table = append(table[:cur], table[cur+1:]...)
		}
		ans := make([]byte, 2*n)
		for i := range ans {
			ans[i] = 'B'
		}
		for _, seat := range table {
			ans[seat] = 'G'
		}

		// 输出，50字符换行
		for i := 0; i < 2*n; i++ {
			fmt.Printf("%c", ans[i])
			if (i+1)%50 == 0 {
				fmt.Println()
			}
		}
		fmt.Println("\n")
	}
}
