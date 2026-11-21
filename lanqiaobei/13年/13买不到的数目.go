/*
小明开了一家糖果店。他别出心裁：把水果糖包成4颗一包和7颗一包的两种。糖果不能拆包卖。
小朋友来买糖的时候，他就用这两种包装来组合。当然有些糖果数目是无法组合出来的，比如要买  10  颗糖。
你可以用计算机测试一下，在这种包装情况下，最大不能买到的数量是17。大于17的任何数字都可以用4和7组合出来。
本题的要求就是在已知两个包装的数量时，求最大不能组合出的数字。
input: 两个正数
output 最大不能买到的糖数
*/
/*
gcd(a, b) = 1, 最大不能组合出来的数字
				ans = a * b - a - b
*/
/*
package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if gcd(a, b) != 1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(a*b - a - b)
}
*/
/*
定义 dp[i] 是否能组合出i颗糖
dp[i] = dp[i-a] || dp[i-b]
dp[0] = true
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	const N = 100000
	// dp[i] 表示能否组合出i颗糖果
	dp := make([]bool, N)
	dp[0] = true
	ans := -1
	for i := 1; i < N; i++ {
		if i >= a && dp[i-a] {
			dp[i] = true
		}
		if i >= b && dp[i-b] {
			dp[i] = true
		}
		if !dp[i] {
			ans = i
		}
	}
	fmt.Println(ans)
}
