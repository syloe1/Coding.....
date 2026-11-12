/*
输入 n 个数字 ai​ ，输出这些数字中第 k 小的数。最小的数是第 0 小。
n <= 5e6 ai < 1e9
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func quickSelect(arr []int, l, r, k int) int {
	if l >= r {
		return arr[l]
	}
	pivot := arr[(l+r)/2]
	i, j := l, r
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	// k 在左边 [l, j]
	if k <= j {
		return quickSelect(arr, l, j, k)
	} else {
		// k是全局下标
		// k 在右边 [i, r]
		return quickSelect(arr, i, r, k)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<25), 1<<25)
	sc.Scan()
	parts := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	sc.Scan()
	numStr := strings.Split(sc.Text(), " ")
	a := make([]int, n)
	for idx := 0; idx < n; idx++ {
		a[idx], _ = strconv.Atoi(numStr[idx])
	}
	ans := quickSelect(a, 0, n-1, k)
	fmt.Println(ans)
}
