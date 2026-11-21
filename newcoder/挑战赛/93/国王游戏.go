//https://ac.nowcoder.com/acm/contest/141878/B
/*
题目B 国王游戏
等级：1国王，0平民，-1囚徒
规则(u,v)，输入保证 u<v
触发条件：轮到u行动时，a[u]=0，a[v]=1
效果：u变成1，v变成-1
行动顺序 n, n-1 ... 1，大编号先行动，小编号后行动。
所有人目标最大化自己最终等级。

n次独立游戏：
第i场：只有i可以选初始a_i ∈ {-1,0,1}，其余初始全部0。
求第i场结束i可以达到的最大最终等级。

图建模：规则(u,v)看作有向边 u→v，表示u可以抢夺v的王位。
建反向邻接表 in[v]：保存所有x，存在边 x→v，也就是所有可以抢v的人。

定义ok[v]：bool，若v拿到国王1，游戏跑完之后v是否可以最终保住1。
ok[v]=true：可以保住；false：就算拿到王位最后一定变囚徒。

递推：v从小到大遍历1~n（因为x<v，x先算完ok[x]）
ok[v] = true
遍历所有 x ∈ in[v]：
    如果 ok[x]==true：存在一个x能抢到王位并且守住，x会抢v
    ok[v]=false break

含义：只要存在一个x，x能抢v并且x抢到后能守住王位，则v守不住王位ok[v]=false。

答案：
ok[i]==true →输出1：i开局选1，所有威胁i的x都守不住王位，x不愿抢夺，i保住国王。
ok[i]==false →输出0：不能选初始1，选初始0；没有人可以把i变成-1，结局0。

答案永远不会输出-1，我们总可以选初始0规避变成囚徒。

样例输入边
1→2,1→5,1→3,2→4,3→4,4→5,2→3
in[1] = []
in[2] = [1]
in[3] = [1,2]
in[4] = [2,3]
in[5] = [1,4]

计算ok
v=1 in为空 ok[1]=true
v=2 x=1 ok[x]=true → ok[2]=false
v=3 x=1 ok[x]=true → ok[3]=false
v=4 x=2 ok=false；x=3 ok=false → ok[4]=true
v=5 x=1 ok=true → ok[5]=false

ok数组 [_,T,F,F,T,F]
输出：
1
0
0
1
0

复杂度 O(n+m)，n,m<=1e5

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 快读，适配1e5规模
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// in[v] 存所有可以抢v的u，也就是边 u->v
	in := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		in[v] = append(in[v], u)
	}

	ok := make([]bool, n+1)
	// 从小到大遍历 1~n，u<v，计算ok[v]的时候ok[u]已经算好
	for v := 1; v <= n; v++ {
		ok[v] = true
		for _, x := range in[v] {
			if ok[x] {
				ok[v] = false
				break
			}
		}
	}

	// 输出答案：ok[i] 则 1，否则 0
	for i := 1; i <= n; i++ {
		if ok[i] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
