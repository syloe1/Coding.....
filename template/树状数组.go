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