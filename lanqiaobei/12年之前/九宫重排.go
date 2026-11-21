/*
如下面第一个图的九宫格中，放着 1~8 的数字卡片，还有一个格子空着。与空格子相邻的格子中的卡片可以移动到空格中。经过若干次移动，可以形成第二个图所示的局面。
我们把第一个图的局面记为：12345678.
把第二个图的局面记为：123.46758
显然是按从上到下，从左到右的顺序记录数字，空格记为句点。
本题目的任务是已知九宫的初态和终态，求最少经过多少步的移动可以到达。如果无论多少步都无法到达，则输出 -1。
## 输入格式
   输入第一行包含九宫的初态，第二行包含九宫的终态。
## 输出格式
   输出最少的步数，如果不存在方案，则输出 -1。
*/
/*
这道题就是经典的 八数码问题（Eight Puzzle），属于图论中的最短路径问题。
把游戏状态看成图中的节点，把一次合法移动看成一条边，然后求从初态到终态的最短路径。
每次取出一个状态：
	1. 找到 . 的位置。
	2. 判断空格能否向上下左右移动。
	3. 交换空格与相邻数字。
	4. 得到一个新状态。
	5. 如果没访问过，就加入队列，距离加一。
还有一个细节：一定要使用 visited，避免状态反复访问。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	state string //当前九宫格状态
	step  int    //从初态走了多少步
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

// bfs 求最少移动步数
func bfs(st, target string) int {
	if st == target {
		return 0
	}
	q := []Node{
		state: st,
		step:  0,
	}
	vs := make(map[string]bool)
	vs[st] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		pos := strings.Index(cur.state, ".")
		x := pos / 3
		y := pos % 3
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx < 0 || nx >= 3 || ny < 0 || ny >= 3 {
				continue
			}
			newpos := nx*3 + ny
			State := []byte(cur.state)
			State[pos], State[newpos] = State[newpos], State[pos]

			nState := string(State)
			if vs[nState] {
				continue
			}
			vs[nState] = true
			if nState == target {
				return cur.step + 1
			}
			// enqueue
			q = append(q, Node{
				state: nState,
				step:  cur.step + 1,
			})
		}
	}
	return -1
}
func main() {
	in := bufio.NewReader(os.Stdin)
	var st, target string
	fmt.Fscan(in, &st, &target)
	fmt.Println(bfs(st, target))
}
