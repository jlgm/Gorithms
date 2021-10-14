type LRUCache struct {
	queue    *list.List
	cache    map[int]*list.Element
	capacity int
}

type pair struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		queue:    list.New(),
		cache:    make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.cache[key] == nil {
		return -1
	}
	elem := this.cache[key]
	ans := elem.Value.(pair)
	this.queue.MoveToFront(elem)
	return ans.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache[key] == nil {
		elem := this.queue.PushFront(pair{key, value})
		this.cache[key] = elem
	} else {
		elem := this.cache[key]
		elem.Value = pair{key, value}
		this.queue.MoveToFront(elem)
	}
	this.CheckCapacity()
}

func (this *LRUCache) CheckCapacity() {
	if this.queue.Len() > this.capacity {
		last := this.queue.Back()
		p := last.Value.(pair)
		this.cache[p.key] = nil
		this.queue.Remove(last)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */