// LRU缓存
/*
想象你有一摞书。
get：把一本书（key）抽出来，放在最上面。
put：放入一本新书。
如果已经有这本书（key），
就把它抽出来放在最上面，并替换它的 value。
（例如把一本书的第二版替换成第三版）
如果没有这本书（key），就放在最上面。
如果超过 capacity 本书，就把最下面的书移除。
> 这就是 LRU（最近最少使用）缓存的通俗比喻。
*/
// 双链表 + 哈希表， 查找O(1), 修改O(1)
type Node struct {
	key, value int
	prev, next *Node
}
type LRUCache struct {
	mp       map[int]*Node
	capacity int
	dummy    *Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.prev = dummy
	dummy.next = dummy
	return LRUCache{
		capacity: capacity,
		dummy:    dummy,
		mp:       make(map[int]*Node),
	}
}
func (c *LRUCache) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
}
func (c *LRUCache) pushFront(x *Node) {
	x.prev = c.dummy
	x.next = c.dummy.next
	x.prev.next = x
	x.next.prev = x
}
func (c *LRUCache) getNode(key int) *Node {
	node := c.mp[key]
	if node == nil {
		return nil
	}
	c.remove(node)
	c.pushFront(node)
	return node
}
func (c *LRUCache) Get(key int) int {
	node := c.getNode(key)
	if node == nil {
		return -1
	}
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	node := c.getNode(key) // getNode 会把对应节点移到链表头部
	if node != nil {       // 有这本书
		node.value = value // 更新 value
		return
	}
	node = &Node{key: key, value: value} // 新书
	c.mp[key] = node
	c.pushFront(node)           // 放到最上面
	if len(c.mp) > c.capacity { // 书太多了
		backNode := c.dummy.prev
		delete(c.mp, backNode.key)
		c.remove(backNode) // 去掉最后一本书
	}
}
