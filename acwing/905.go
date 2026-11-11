/*
每次想让点盖住尽可能多后面的区间， 点放越靠右，越容易蹭到后面和它重叠的区间。
1. 把所有区间，**按照区间的右端点从小到大排序**
2. 记录上一个选中的点 `last`（一开始设成很小很小，相当于还没选点）
3. 逐个遍历区间：
  - 如果**这个区间的左端点 > last**：说明上一个选的点已经不在这个区间里了，必须新选一个点
  - 新点就选**当前区间的右端点**，更新 last，计数 +1
  - 否则：上一个点还在本区间内，不用管
*/
package main

import (
	"os"
	"sort"
)

const bufSize = 1 << 20 // 1MB buffer，不够就改成 1<<21

var buf [bufSize]byte
var ptr int

func read() int {
	sign := 1
	for ; buf[ptr] < '0' || buf[ptr] > '9'; ptr++ {
		if buf[ptr] == '-' {
			sign = -1
			ptr++
			break
		}
	}
	res := 0
	for ; buf[ptr] >= '0' && buf[ptr] <= '9'; ptr++ {
		res = res*10 + int(buf[ptr]-'0')
	}
	return res * sign
}

type seg struct {
	l, r int
}

func main() {
	n, _ := os.Stdin.Read(buf[:])
	_ = n

	N := read()
	segs := make([]seg, N)
	for i := range segs {
		segs[i].l = read()
		segs[i].r = read()
	}

	sort.Slice(segs, func(i, j int) bool {
		return segs[i].r < segs[j].r
	})

	ans, last := 0, -1<<60
	for _, s := range segs {
		if s.l > last {
			ans++
			last = s.r
		}
	}
	println(ans)
}
