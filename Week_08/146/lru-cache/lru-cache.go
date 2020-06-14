package lru_cache

// https://leetcode.com/problems/lru-cache/

// 执行用时 : 120 ms, 在所有 Go 提交中击败了 99.56% 的用户
// 内存消耗 : 15.3 MB, 在所有 Go 提交中击败了 90.91% 的用户
type LRUCache struct {
	cache      map[int]*linkNode
	head, tail *linkNode
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head, tail := &linkNode{}, &linkNode{}
	head.next, tail.pre = tail, head
	return LRUCache{
		cache:    map[int]*linkNode{},
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	cache := this.cache
	pre := this.tail.pre
	if node, ok := cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	if len(cache) == this.capacity {
		delete(cache, pre.key)
		this.removeNode(pre)
	}
	node := &linkNode{key: key, val: value}
	cache[key] = node
	this.addNode(node)
}

func (this *LRUCache) addNode(node *linkNode) {
	head := this.head
	node.next = head.next
	node.pre = head
	head.next.pre = node
	head.next = node
}

func (this *LRUCache) removeNode(node *linkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *linkNode) {
	this.removeNode(node)
	this.addNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */