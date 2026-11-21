// https://ac.nowcoder.com/acm/contest/143407/A
/*
max sigma(R[i] - L[i] + 1) ^ 2 转换成
	min sigma(R[i] * L[i]) 然后发现应该让大的左端点匹配小的合法右端点。
 树状数组维护右端点出现次数，支持删除和第 k 小查询。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod int64 = 998244353

type Fenwick struct {
	n   int
	bit []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		n:   n,
		bit: make([]int, n+1),
	}
}
func (f *Fenwick) Add(i, v int) {
	for i <= f.n {
		f.bit[i] += v
		i += i & -i
	}
}

// 查询前缀和
func (f *Fenwick) Sum(i int) int {
	res := 0
	for i > 0 {
		res += f.bit[i]
		i -= i & -i
	}
	return res
}
func (f *Fenwick) Kth(k int) int {
	pos := 0
	step := 1
	for step<<1 <= f.n {
		step <<= 1
	}
	for step > 0 {
		nxt := pos + step
		if nxt <= f.n && f.bit[nxt] < k {
			pos = nxt
			k -= f.bit[nxt]
		}
		step >>= 1
	}
	return pos + 1
}
func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	L := make([]int64, n)
	R := make([]int64, n)
	rs := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &L[i], &R[i])
		rs[i] = R[i]
	}
	sort.Slice(L, func(i, j int) bool {
		return L[i] > L[j]
	})
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	un := make([]int64, 0, n)
	for _, x := range rs {
		if len(un) == 0 || un[len(un)-1] != x {
			un = append(un, x)
		}
	}
	m := len(un)
	fw := NewFenwick(m)
	for _, r := range R {
		idx := sort.Search(len(un), func(i int) bool {
			return un[i] >= r
		}) + 1
		fw.Add(idx, 1)
	}
	var ans int64
	for _, l := range L {
		k := sort.Search(len(un), func(i int) bool {
			return un[i] >= l
		})
		cnt := fw.Sum(k)
		idx := fw.Kth(cnt + 1)
		r := un[idx-1]
		fw.Add(idx, -1)
		ql := r - l + 1
		v := ql % mod
		ans = (ans + v*v%mod) % mod
	}
	fmt.Println(ans)
}
