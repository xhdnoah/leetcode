package main

type LRUCache struct {
	head, tail *Node         // 双向链表的伪头部和伪尾部
	Keys       map[int]*Node // 哈希表
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

// Input:
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output: [null, null, null, 1, null, -1, null, -1, 3, 4]
// Explanation: 最近最少使用淘汰 Least Recently Used
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4
// LRU 更新和插入发生在链表首部，删除尾部发生在插入时超出容量
// 双向链表按照被使用的顺序存储键值对，头部数据最常使用而尾部则最近最少使用
// 哈希表 普通的哈希映射 HashMap 通过缓存数据的键映射到其在双向链表中的位置
// 在双向链表的实现中，使用伪头部、伪尾部标记界限 免去在添加或删除节点时检查相邻节点是否存在
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Keys: make(map[int]*Node),
		head: &Node{},
		tail: &Node{},
		Cap:  capacity,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

// 添加新结点至双向链表表头
func (this *LRUCache) addNode(node *Node) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

// 删除双向链表任意结点
func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next // 断开与前结点的连接
	node.Next.Prev = node.Prev // 断开与后结点的连接
}

func (this *LRUCache) addRecently(key, val int) {
	node := &Node{Key: key, Val: val}
	this.Keys[key] = node
	this.addNode(node)
}

func (this *LRUCache) makeRecently(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) removeLeastRecently() {
	node := this.tail.Prev
	delete(this.Keys, node.Key)
	this.removeNode(node)
}

// Get Put 必须以 O(1) 的平均时间复杂度运行 使用 HashMap
// 在 map 中直接读取双向链表的结点, 存在就移动到双向链表表头并返回 value
func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.makeRecently(node) // 将结点提升至最近使用
		return node.Val
	}
	return -1
}

// 查询 map 存在 key 就更新 value 并移动到表头，不存在就新建
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value // 更新 map
		this.makeRecently(node)
		return
	}
	// 维护双向链表的 cap 若超出则淘汰最后结点
	if len(this.Keys) == this.Cap {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}
