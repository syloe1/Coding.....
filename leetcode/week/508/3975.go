/*
给你一个二维整数数组 occupiedIntervals，其中 occupiedIntervals[i] = [starti, endi] 表示你处于忙碌状态的一个时间区间。每个区间从 starti 开始，到 endi 结束，并且 包含 两个端点。这些区间可能会 重叠。
此外，另给你两个整数 freeStart 和 freeEnd，它们定义了一个你空闲的时间区间。该空闲区间从 freeStart 开始，到 freeEnd 结束，并且 包含 两个端点。
你的任务是先将所有重叠或相接的忙碌区间 合并 ，然后从合并后的忙碌区间中 移除 空闲区间内的 所有 整数点。
如果第二个区间正好从第一个区间结束后的下一个位置开始，则称这两个区间相接。例如，[1, 1] 和 [2, 2] 相接，应合并为 [1, 2]。
返回按 升序 排列的 剩余 忙碌区间。返回的区间必须 互不重叠 ，并且区间数量应尽可能 最少 。如果没有剩余的忙碌整数点，则返回 空列表 。
length <= 5e4
oc[i].len == 2
st <= ed <= 1e9
freeSt <= freeEd <= 1e9
*/
/*

整个算法就是：
1. 对 occupiedIntervals 按照左端点升序排序。
2. 遍历区间，合并所有重叠或相接的区间。
3. 遍历合并后的区间。
4. 如果当前区间与空闲区间没有交集，直接加入答案。
5. 如果有交集，分别判断是否存在左侧剩余部分和右侧剩余部分。
6. 将剩余部分加入答案并返回。
*/
func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) [][]int {
	// 第一步：按照左端点升序排序
	sort.Slice(occupiedIntervals, func(i, j int) bool {
		return occupiedIntervals[i][0] < occupiedIntervals[j][0]
	})

	// 第二步：合并重叠或相接的区间
	merged := make([][]int, 0)

	for _, interval := range occupiedIntervals {
		start, end := interval[0], interval[1]

		if len(merged) == 0 || start > merged[len(merged)-1][1]+1 {
			// 不能合并，直接加入
			merged = append(merged, []int{start, end})
		} else {
			// 可以合并，更新右端点
			last := merged[len(merged)-1]
			if end > last[1] {
				last[1] = end
			}
		}
	}

	// 第三步：从合并后的区间中删除空闲区间
	ans := make([][]int, 0)

	for _, interval := range merged {
		start, end := interval[0], interval[1]

		// 没有交集，原样保留
		if end < freeStart || start > freeEnd {
			ans = append(ans, interval)
			continue
		}

		// 保留空闲区间左边的部分
		if start < freeStart {
			ans = append(ans, []int{start, freeStart - 1})
		}

		// 保留空闲区间右边的部分
		if end > freeEnd {
			ans = append(ans, []int{freeEnd + 1, end})
		}
	}

	return ans
}