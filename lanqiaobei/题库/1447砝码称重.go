// https://www.lanqiao.cn/problems/1447/learning/?page=1&first_category_id=1&name=%E7%A0%9D%E7%A0%81%E7%A7%B0%E9%87%8D
/*
【问题描述】
你有一架天平和 N 个砝码，这 N 个砝码重量依次是 W1,W2,…,WN。
请你计算一共可以称出多少种不同的重量。注意砝码可以放在天平两边。

【输入格式】
输入的第一行包含一个整数 N。
第二行包含 N 个整数：W1,W2,W3,…,WN。

【输出格式】
输出一个整数表示答案。

【数据范围】
对于所有评测用例，1 ≤ N ≤ 100，N 个砝码总重不超过 10000。

*/

/*
解题思路：天平砝码（砝码可放两边）
1. 分析选择
每个砝码有3种选择：
① 不使用这个砝码
② 放到砝码侧（物体对面）：重量 +w
③ 放到物体同侧：重量 -w
普通01背包只有选/不选，本题是带负数的01背包。

2. 负数下标问题
直接存差值会出现负数，数组下标不能为负。
设所有砝码总重量为sum，偏移量 offset = sum。
真实差值 s 对应数组下标：idx = s + offset
s=0 → idx=sum
s=-5 → idx=sum-5
s=5 → idx=sum+5

3. dp数组定义
dp是bool数组，dp[idx] = true 代表可以凑出差值 (idx - offset)
初始状态：dp[offset] = true，代表不用任何砝码，差值为0。

4. 状态转移（01背包滚动更新）
遍历每一个砝码w：
倒序遍历背包，防止同一个砝码重复选取
如果dp[j]可达：
    dp[j + w] = true
    dp[j - w] = true
不选的情况自动保留，不需要额外操作。

5. 统计答案
只统计真实差值 > 0 的情况（也就是 j > offset），差值为0不算（不能称出物体）。
统计有多少个位置dp[j]为true，就是可以称出的不同重量数量。

6. 边界与复杂度
数据范围：N<=100，总重量sum<=10000
偏移后数组最大长度 2*sum，时间复杂度 O(N*sum)，可以通过。

举例：N=2，砝码[1,3]，sum=4，offset=4
可达正数：1,2,3,4，答案4。

*/
/*
dp[i][s] 只考虑前i个砝码， 能不能凑出差值等于s
dp[i][s] = true, 可行
dp[i][s] = false, 不行
s 是什么 ？   sum是另一侧砝码之和
	物体x和砝码t不在同一侧  x  + (-t) = sum , x = sum + t  贡献+t
	物体x和砝码t在同一侧   x + t = sum, x = sum - t 贡献-t
初始
dp[0][0] = true
dp[0[s]] = false ,
不用第i个砝码 	dp[i][Sold] = dp[i-1][Sold]
砝码和物体对面   物体和另一边差值变大 S = Sold + wi
			dp[i][Sold+w[i] || dp[i-1][Sold]
砝码和物体同一侧 差值变小   S = Sold - wi
			dp[i][Sold-w[i]] || dp[i-1][Sold]
dp[i][s] = dp[i-1][s] || dp[i-1][s-w[i]] || dp[i-1][s +w[i]]
s有负数，增加一个非偏移， 设砝码总和为F
真实差值=数组下标 ， idx = s + F
					s = idx - F
不选 idx = idx_old, 放对面  s=sold+w[i] idx=idx_old+w[i]
					放同侧  s=sold-w[i], idx = idx_old-w[i]
		s in [-F, F] => idx in [0， 2F]
dp[i][idx] 前i个砝码 能不能凑出差值s=idx-F
	dp[i][idx] = dp[i-1][idx] || dp[i-1][idx-w[i]] || dp[i-1][idx+w[i]]
初始dp[0][F] = true
 dp[0][0] 下标=0+F
*/
package main

import "fmt"

// TwoDimBalance 天平砝码二维DP版本
// n:砝码数量, w:砝码数组(下标从1开始)
// 返回值：可以称出的不同重量的种类数
func TwoDimBalance(n int, w []int) int {
	F := 0
	for i := 1; i <= n; i++ {
		F += w[i]
	}
	offset := F
	// dp[i][idx] 前i个砝码，能否凑出差值s = idx - offset
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, 2*F+1)
	}
	dp[0][offset] = true
	for i := 1; i <= n; i++ {
		wi := w[i]
		for idx := 0; idx <= 2*F; idx++ {
			dp[i][idx] = dp[i-1][idx]
			// 放同一侧
			if idx-wi >= 0 {
				dp[i][idx] = dp[i][idx] || dp[i-1][idx-wi]
			}
			if idx+wi <= 2*F {
				dp[i][idx] = dp[i][idx] || dp[i-1][idx+wi]
			}
		}
	}
	ans := 0 // 统计s > 0
	for idx := offset + 1; idx <= 2*F; idx++ {
		if dp[n][idx] {
			ans++
		}
	}
	return ans
}

/*
优化
一旦修改了某个 `dp[idx]`，后面不管往左、往右读，**读到的就可能是本轮 i 刚更新的新值，不是 i-1 的旧值 → 砝码重复使用，变成完全背包，直接错**
每次处理一个砝码之前，先把 i-1 的完整 dp 数组复制一份备份（叫 oldDp），更新全部从 oldDp 读，写在原 dp 上。
dp[idx] 当前处理完若干砝码，能不能抽出差值s=idx-F
dp[F] = true
dp[idx] = oldDp[idx] ||
          (idx-wi >=0 && oldDp[idx-wi]) ||
          (idx+wi <=2*F && oldDp[idx+wi])
 oldDp[idx] 不选 oldDp[idx-wi] 同侧 oldDp[idx+wi] 放对面
*/
// OneDimBalance 天平砝码一维DP版本
// n:砝码数量, w:砝码数组(下标从1开始)
// 返回值：可以称出的不同重量的种类数
func OneDimBalance(n int, w []int) int {
	F := 0
	for i := 1; i <= n; i++ {
		F += w[i]
	}
	offset := F
	size := 2*F + 1
	dp := make([]bool, size)
	dp[offset] = true
	for i := 1; i <= n; i++ {
		wi := w[i]
		// 备份i-1层的全部状态
		oldDp := make([]bool, size)
		copy(oldDp, dp)
		for idx := 0; idx < size; idx++ {
			ok := oldDp[idx]
			// 和物体同侧
			if idx-wi >= 0 {
				ok = ok || oldDp[idx-wi]
			}
			if idx+wi < size {
				ok = ok || oldDp[idx+wi]
			}
			dp[idx] = ok
		}
	}
	ans := 0
	for idx := offset + 1; idx < size; idx++ {
		if dp[idx] {
			ans++
		}
	}
	return ans
}
func main() {
	var n int
	fmt.Scan(&n)
	w := make([]int, n+1) // 开 n+1，下标1开始
	for i := 1; i <= n; i++ {
		fmt.Scan(&w[i])
	}
	fmt.Println(OneDimBalance(n, w))
}
