package other

type LRUCache struct {
	size, cap  int
	head, tail *Node
	cache      map[int]*Node
}

type Node struct {
	key, value int
	pre, next  *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:  0,
		cap:   capacity,
		head:  initNode(0, 0),
		tail:  initNode(0, 0),
		cache: make(map[int]*Node, capacity),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func initNode(key, value int) *Node {
	return &Node{key: key, value: value}
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node != nil {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node == nil {
		this.size++
		if this.size > this.cap {
			node = this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}

		node = initNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head

	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
