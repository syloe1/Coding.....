package main

import "fmt"

func main() {
	var n, m int
	// n：数值上限，m：一共有多少个数字
	fmt.Scan(&n, &m)

	// 计数数组，下标0不用，1~n
	cnt := make([]int, n+1)

	// 读m个数，统计次数
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		cnt[tmp]++
	}

	// 从小到大展开输出
	for i := 1; i <= n; i++ {
		// 这个数字出现多少次，就打印多少次
		for j := 0; j < cnt[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
