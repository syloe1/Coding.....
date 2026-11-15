/*
给你一个整数数组 nums。
如果一个子数组中所有元素的 按位与 运算结果等于该子数组的 最小值 ，则称该子数组为 一致子数组 。
返回 nums 中 一致子数组 的个数。
子数组 是数组中一个连续、非空的元素序列。
n < 2e5  nums[i] < 1e5
*/
package main

import "sort"

func countConsistent(nums []int) int64 {
	type pair struct {
		val int // 当前分组的and值
		l   int // 当前分组最小的左端点
	}
	//pos[x] x在nums中出现的所有位置
	pos := make(map[int][]int)
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}
	var ans int64
	vec := make([]pair, 0, 21)
	for i, x := range nums {
		vec = append(vec, pair{x, i})
		nxt := make([]pair, 0, 21)
		for _, p := range vec {
			p.val &= x
			if len(nxt) == 0 ||
				nxt[len(nxt)-1].val != p.val {
				nxt = append(nxt, p)
			}
		}
		vec = nxt
		// 统计每个AND分组的贡献
		for j, p := range vec {
			l := p.l
			r := i + 1
			// 下一个分组的起点， 就是当前分组的有边界
			if j+1 < len(vec) {
				r = vec[j+1].l
			}
			x := p.val
			// 找最后一个小于r的x的出现位置
			positions := pos[x]
			idx := sort.SearchInts(positions, r)
			if idx > 0 {
				last := positions[idx-1]
				if last >= l {
					ans += int64(last - l + 1)
				}
			}
		}
	}
	return ans
}
