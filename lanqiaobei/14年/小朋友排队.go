/*
n  个小朋友站成一排。现在要把他们按身高从低到高的顺序排列，但是每次只能交换位置相邻的两个小朋友。
每个小朋友都有一个不高兴的程度。开始的时候，所有小朋友的不高兴程度都是0。
如果某个小朋友第一次被要求交换，则他的不高兴程度增加1，如果第二次要求他交换，则他的不高兴程度增加2（即不高兴程度为3），依次类推。当要求某个小朋友第k次交换时，他的不高兴程度增加k。
请问，要让所有小朋友按从低到高排队，他们的不高兴程度之和最小是多少。
如果有两个小朋友身高一样，则他们谁站在谁前面是没有关系的。
样例说明
首先交换身高为3和2的小朋友，再交换身高为3和1的小朋友，再交换身高为2和1的小朋友，每个小朋友的不高兴程度都是3，总和为9。
输入的第一行包含一个整数n，表示小朋友的个数。
第二行包含  n  个整数  H1  H2  …  Hn，分别表示每个小朋友的身高。

数据规模和约定
对于100%的数据，1< =n< =100000，0< =Hi< =1000000。
输出一行，包含一个整数，表示小朋友的不高兴程度和的最小值。
*/
/*
遇到相邻交换排序问题，不要急着模拟交换过程。

先问自己三个问题：
1. 哪些元素之间必须交换？通常与逆序对有关。
2. 每个元素最终需要参与多少次交换？
3. 每次交换的代价是否与参与次数有关？
小朋友换k次， 不高兴cheng程度 1 + ... + k
			k(k*1)/2
sigma(i form 1 to n) k*(k*1)/2
可以发现，每一对身高不符合顺序的小朋友，最终都必须交换一次。这就是逆序对。
eg: 312 逆序对2， 3交换2次, 1交换1次，2交换1次
总交换次数 = 2 * 逆序对
原位置为 \(i\) 的小朋友：
- 左边比他高、需要移动到他后面的人，会和他交换。
- 右边比他矮、需要移动到他前面的人，也会和他交换。
*/ // 离散化 + 树状数组 nlogn

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
普通数组 修改 O1 查询前缀和On
树状数组 修改 Ologn 查询O(logn)
不直接存储每个位置的前缀和，而是让每个位置负责一段区间的和
lowbit(i) = i & -i 取出二进制中最低位1所代表的值
tree[i] 负责的区间 [i-lowbit(i)+1, i]
树状数组不需要把 1+2+3+4+5 全部加起来，而是：[sum(5)=tree[4]+tree[5]]
*/
type BIT struct {
	tree []int
	n    int
}

/*
Add 修改 下标 i += lowbit(i)
Query查询 下标  i -= lowbit(i)
前面有多少数比5大 ？？ [前面元素总数-前面小于等于5的元素数量]

	小于5的数量就是 bit.Query(rank[5])

l[i] = i - bit.Query(rank[i])
树状数组通过下标来维护频率
离散化的意义：

	不改变元素之间的大小关系，把很大的数值映射成较小的连续整数。
	val太大，数组不好用val做下标
*/
func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1),
		n:    n,
	}
}

// 单点修改 包含v的区间都要修改
func (b *BIT) Add(i, v int) {
	for i <= b.n {
		b.tree[i] += v
		// i & -i 是lowbit
		i += i & -i
	}
}

// 统计排名小于等于x的元素数量
func (b *BIT) Query(i int) int {
	res := 0
	for i > 0 {
		res += b.tree[i]
		i -= i & -i
	}
	return res
}

type Person struct {
	h  int
	id int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]Person, n)
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].h)
		a[i].id = i
		vals[i] = a[i].h
	}
	// 离散化
	sort.Ints(vals)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[i] = sort.SearchInts(vals, a[i].h) + 1
	}
	l := make([]int64, n)
	bit := NewBIT(n)
	for i := 0; i < n; i++ {
		idx := bit.Query(rank[i])
		l[i] = int64(i - idx)
		bit.Add(rank[i], 1)
	}
	r := make([]int64, n)
	bit = NewBIT(n)
	// 从右向左遍历， 直接query， -1减去等于的那个
	for i := n - 1; i >= 0; i-- {
		r[i] = int64(bit.Query(rank[i] - 1))
		bit.Add(rank[i], 1)
	}
	var ans int64
	for i := 0; i < n; i++ {
		k := l[i] + r[i]
		ans += k * (k + 1) / 2
	}
	fmt.Fprintln(out, ans)
}
