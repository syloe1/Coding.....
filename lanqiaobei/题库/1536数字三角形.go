// https://www.lanqiao.cn/problems/1536/learning/?page=1&first_category_id=1&name=%E6%95%B0%E5%AD%97%E4%B8%89%E8%A7%92%E5%BD%A2
/*
【问题描述】
上图给出了一个数字三角形。从三角形的顶部到底部有很多条路径。对于每条路径，把路径上面的数加起来可以得到一个和，你的任务就是找到最大的和。
路径上的每一步只能从一个数走到下一层和它最近的左边的那个数或者右边的那个数。此外，向左下走的次数与向右下走的次数相差不能超过 1。

【输入格式】
输入的第一行包含一个整数 N (1 ≤ N ≤ 100)，表示三角形的行数。
下面的 N 行给出数字三角形。数字三角形上的数都是 0 至 100 之间的整数。

【输出格式】
输出一个整数表示答案。

*/
// 没有步数限制的
/*
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			fmt.Scan(&mat[i][j])
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = mat[0][0]
	for i := 1; i < n; i++ {
		// 最左
		dp[i][0] = dp[i-1][0] + mat[i][0]
		// 最右
		dp[i][i] = dp[i-1][i-1] + mat[i][i]
		// 中间
		for j := 1; j < i; j++ {
			if dp[i-1][j] > dp[i-1][j-1] {
				dp[i][j] = dp[i-1][j] + mat[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + mat[i][j]
			}
		}
	}
	ans := dp[n-1][0]
	for j := 1; j < n; j++ {
		if dp[n-1][j] > ans {
			ans = dp[n-1][j]
		}
	}
	fmt.Println(ans)
}
自底向上版本（推荐，代码短，不用处理边界）
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			t := max(mat[i+1][j], mat[i+1][j+1])
			mat[i][j] += t
		}
	}
	fmt.Println(mat[0][0])
*/
// 带步数限制
/*
L 左下次数， R 右下次数
L + R = N - 1, |L - R| <= 1
k = n-1 偶数  L = R， 左下j不变， 右下j+1 终点列j = k / 2
		奇数 R = L + 1 终点列就是 j = (k-1)/2 or (k+1) /2
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = a[0][0]

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				// 只能从右上（上一层j=0）过来
				dp[i][j] = dp[i-1][j] + a[i][j]
			} else if j == i {
				// 只能从左上（上一层j-1）过来
				dp[i][j] = dp[i-1][j-1] + a[i][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]) + a[i][j]
			}
		}
	}
	k := n - 1
	var ans int
	if k%2 == 0 {
		ans = dp[n-1][k/2]
	} else {
		p1 := k / 2
		p2 := p1 + 1
		ans = max(dp[n-1][p1], dp[n-1][p2])
	}
	fmt.Println(ans)

}

/*
5
7
3 8
8 1 0
2 7 4 4
4 5 2 6 5
*/
