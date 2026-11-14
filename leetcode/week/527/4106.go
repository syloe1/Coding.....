/*
给你两个长度相同的字符串 s1 和 s2。

不匹配片段 是一个 极大连续 下标区间，其中每个下标处 s1 和 s2 的字符都不相同。

返回一个二维数组，其中包含每个不匹配片段对应的 [sub1, sub2]，sub1 和 sub2 分别是 s1 和 s2 中对应的 子字符串。请按片段起始下标的 递增 顺序返回。如果没有不匹配片段，则返回一个空数组。
示例 1：

输入： s1 = "earth", s2 = "empty"

输出： [["ar","mp"],["h","y"]]

解释：

下标 1 和 2 处的字符不同，形成不匹配片段 ["ar","mp"]。
下标 4 处的字符不同，形成不.匹配片段 ["h","y"]。
1 <= s1.length == s2.length <= 100
s1 和 s2 仅由小写英文字母组成。
设定一个变量 start = -1，用来记录当前不匹配片段的起始位置。

遍历字符串的每个字符：

如果字符不同 (s1[i] != s2[i])：如果 start == -1，说明这是一个新片段的开始，记录 start = i。

如果字符相同 (s1[i] == s2[i])：如果 start != -1，说明上一个不匹配片段结束了。此时截取 s1[start:i] 和 s2[start:i] 放入结果，并把 start 重置为 -1。

边界处理：遍历结束后，如果 start != -1，说明字符串是以不匹配片段结尾的，需要再截取一次。
*/
func findMismatchedSegments(s1 string, s2 string) [][]string {
	res := make([][]string, 0)
	st := -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if st == -1 {
				st = i
			}
		} else {
			if st != -1 {
				res = append(res, []string{s1[st:i], s2[st:i]})
				st = -1
			}
		}
	}
	if st != -1 {
		res = append(res, []string{s1[st:], s2[st:]})
	}
	return res
}