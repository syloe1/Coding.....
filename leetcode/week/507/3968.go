/*
给你一个由字符 'U'、'D'、'L'、'R' 和 '_' 组成的字符串 moves。

从原点 (0, 0) 出发，每个字符表示二维平面上的一次移动：

'U'：向上移动 1 个单位。
'D'：向下移动 1 个单位。
'L'：向左移动 1 个单位。
'R'：向右移动 1 个单位。
'_'：可以独立地替换为 'U'、'D'、'L' 或 'R' 中的任意一个字符。
返回执行完所有移动后，能够达到的距离原点的 最大曼哈顿距离 。

两点 (x1, y1) 和 (x2, y2) 之间的 曼哈顿距离 为 |x1 - x2| + |y1 - y2|。
move.length () <= 1e5
*/
/*
曼哈顿距离
|x|+|y|=max(x+y,x-y,-x+y,-x-y)
先统计所有固定方向的移动，计算它们最终产生的位移，再让所有 _ 尽可能远离原点。
*/
// 如果枚举_下标4种走法， 如果10个， 每个4种， 4**10种
func maxDistance(moves string) int {
	x, y, free := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			x--
		case 'R':
			x++
		case 'D':
			y--
		case 'U':
			y++
		default:
			free++
		}
	}
	return abs(x) + abs(y) + free
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}