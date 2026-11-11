/*

给定一个按照升序排列的长度为 n
 的整数数组，以及 q
 个查询。

对于每个查询，返回一个元素 k
 的起始位置和终止位置（位置从 0
 开始计数）。

如果数组中不存在该元素，则返回 -1 -1。

输入格式
第一行包含整数 n
 和 q
，表示数组长度和询问个数。

第二行包含 n
 个整数（均在 1∼10000
 范围内），表示完整数组。

接下来 q
 行，每行包含一个整数 k
，表示一个询问元素。

输出格式
共 q
 行，每行包含两个整数，表示所求元素的起始位置和终止位置。

如果数组中不存在该元素，则返回 -1 -1。

数据范围
1≤n≤100000

1≤q≤10000

1≤k≤10000
*/
// [l, r) 左闭右开
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(firstLine[0])
	q, _ := strconv.Atoi(firstLine[1])

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}

	for i := 0; i < q; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		// 找左边界：第一个 >= k
		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if a[mid] >= k {
				r = mid
			} else {
				l = mid + 1
			}
		}
		left := l

		// 找右边界：最后一个 <= k
		l, r = 0, n-1
		for l < r {
			mid := (l + r + 1) / 2 // 向上取整！防死循环
			if a[mid] <= k {
				l = mid
			} else {
				r = mid - 1
			}
		}
		right := l

		if a[left] != k {
			fmt.Println("-1 -1")
		} else {
			fmt.Println(left, right)
		}
	}
}
