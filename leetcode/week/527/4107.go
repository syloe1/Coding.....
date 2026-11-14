/*
给你两个长度均为 n 的整数数组 nums 和 target。

对于每个下标 i，考虑将原数组的 后缀 nums[i..n - 1] 按 非递增顺序 排序。

返回使排序后的数组等于 target 的下标 i 的数量。

数组的 后缀 是从任意下标开始并延伸至数组末尾的 子数组。
1 <= n == nums.length == target.length <= 1e5
1 <= nums[i], target[i] <= 1e5​​​​​​​​​​​​​​



候选 i 一共有 n 个：0, 1, 2, ..., n-1

每个 i 都要单独试一次

每次操作都是基于原数组 nums，不是在上一次结果上改

操作规则：

前缀 nums[0..i-1] 不动
后缀 nums[i..n-1] 按非递增（从大到小）排序
拼成新数组
如果新数组等于 target，这个 i 就算一个答案

返回满足条件的 i 的数量
*/
/*
不要真的对每个 i 排序后缀，而是判断三个必要条件是否同时成立。

对于某个下标 i，操作后要等于 target，必须且只需满足：

前缀不变：
nums[0..i-1] == target[0..i-1]

后缀多重集相同：
nums[i..n-1] 和 target[i..n-1] 包含的元素种类及次数完全一样。
因为排序不改变元素集合。

目标后缀本身必须是非递增的：
target[i..n-1] 从大到小排列。
因为排序后的后缀一定是非递增的。

这三个条件合起来就充分了：
目标后缀是非递增的，且和 nums 后缀元素完全一样，那么把 nums 后缀降序排序后，结果必然就是 target 后缀。
*/
/*
整体元素出现次数必须相同
对 nums 的某个后缀排序，不会改变整个数组 nums 里每个数出现的次数。

所以如果操作后要等于 target，那么一开始：nums 中每个数的出现次数 == target 中每个数的出现次数
前缀不能变 nums[0..i-1] == target[0..i-1]
设 maxI 是第一个满足 nums[j] != target[j] 的下标 j。

如果 i > maxI，那么前缀里就包含了 maxI 这个不同的位置，前缀不可能相等。

所以必须 i <= maxI。

如果整个数组一开始就完全相同，则没有不同位置，令 maxI = n-1，所有 i 都允许。

所以 maxI 是 i 的上界。
排序后的后缀一定非递增，所以 target 后缀也必须非递增
1 5 6 4 6  3maxi
1 5 6 6 4
 k = 2 target[1] < targe[2] 后缀不能包含上升的
 i >= k
*/
package main

func countSuffix(nums []int, target []int) int {
	// 大前提：每种元素的出现次数都相等
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}
	for _, x := range target {
		cnt[x]--
	}
	for _, v := range cnt {
		if v != 0 {
			return 0
		}
	}

	n := len(nums)
	//设 maxI 是第一个满足 nums[i] != target[i] 的下标 i。
	maxI := n - 1
	for i, x := range nums {
		if x != target[i] {
			maxI = i
			break
		}
	}

	minI := 0
	//最靠右的上升对的后一个位置
	for i := n - 1; i > 0; i-- {
		if target[i-1] < target[i] {
			minI = i
			break
		}
	}
	/*
	   为什么三个条件合起来就充分？
	   假设：

	   nums 和 target 整体多重集相同；

	   i <= maxI，所以前缀 nums[0..i-1] == target[0..i-1]；

	   i >= minI，所以 target[i..n-1] 非递增。

	   由 1 和 2：

	   整体多重集相同，前缀又相同，那么剩下的后缀多重集也必然相同：

	   text
	   nums[i..n-1] 的多重集 == target[i..n-1] 的多重集
	   由 3：

	   target[i..n-1] 本身已经是非递增的。

	   现在把 nums[i..n-1] 按非递增排序，得到的是这个多重集的唯一降序排列。而 target[i..n-1] 也是同一个多重集，并且已经非递增，所以排序结果必然正好等于 target[i..n-1]。

	   前缀又相等，所以整个数组等于 target。

	   所以答案就是区间里的整数个数
	   满足条件的 i 是：

	   text
	   minI <= i <= maxI
	   数量为：

	   text
	   maxI - minI + 1
	   如果这个数小于 0，说明没有合法的 i，答案取 0。
	*/
	return max(maxI-minI+1, 0)
}
