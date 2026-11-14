/*
有 n 个房间，房间按从 0 到 n - 1 编号。最初，除 0 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。

当你进入一个房间，你可能会在里面找到一套 不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。

给你一个数组 rooms 其中 rooms[i] 是你进入 i 号房间可以获得的钥匙集合。如果能进入 所有 房间返回 true，否则返回 false。
*/
//把房间看成节点： rooms[i] 里的每个钥匙 x，相当于一条有向边 i -> x
// 有向图从起点0个的可达性

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vs := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		vs[u] = true
		for _, v := range rooms[u] {
			if !vs[v] {
				dfs(v)
			}
		}
	}
	dfs(0)
	for _, ok := range vs {
		if !ok {
			return false
		}
	}
	return true
}