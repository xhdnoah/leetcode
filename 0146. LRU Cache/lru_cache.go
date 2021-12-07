package main

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
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
// Explanation: 最近最少使用淘汰
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
// LRU 更新和插入新页面都发生在链表首，删除页面都发生在链表尾
func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

// 在 map 中直接读取双向链表的结点, 存在就移动到双向链表表头并返回 value
func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

// 查询 map 存在 key 就更新 value 并移动到表头，不存在就新建
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	// 最后需维护双向链表的 cap 若超出则淘汰最后结点
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head // 指向当前表头
	if this.head != nil {
		this.head.Prev = node // 表头反指
	}
	this.head = node      // 被赋为新表头
	if this.tail == nil { // 链表初始化时
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next // head 新指
		if node.Next != nil {
			node.Next.Prev = nil // 断开连接
		}
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev // tail 新指
		if node.Prev != nil {
			node.Prev.Next = nil // 断开连接
		}
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next // 断开与前结点的连接
	node.Next.Prev = node.Prev // 断开与后结点的连接
}
